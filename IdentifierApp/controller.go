package identifierapp

import (
	"fmt"
	"net/http"

	"github.com/sepehrkhajeh/panda01/dbconnection"
	"github.com/sepehrkhajeh/panda01/repositories"
	"github.com/sepehrkhajeh/panda01/schemas"

	"github.com/labstack/echo/v4"
)

func CreateIdentifier(repo *repositories.IdentifierRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Parse request body
		req := new(IdentifierCreateRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
		}
		ctx := c.Request().Context()

		// Connect to domain repository
		dom, err := dbconnection.ConnectToRepo("domains")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "server error"})
		}
		if dom == nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "domain repository not found"})
		}

		domainRepo, ok := dom.(*repositories.DomainRepository)
		if !ok {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "invalid repository type"})
		}

		d, err := domainRepo.GetByID(ctx, req.DomainID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error retrieving domain information"})
		}
		if d == nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "domain not found"})
		}

		fullIdentifier := fmt.Sprintf("%s@%s", req.Name, d.Domain)

		// Create new identifier
		newIdentifier := schemas.IdentifierSchema{
			Name:           req.Name,
			Pubkey:         req.Pubkey,
			DomainID:       req.DomainID,
			FullIdentifier: fullIdentifier,
		}

		identifier, err := repo.Add(ctx, newIdentifier)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error creating identifier"})
		}

		return c.JSON(http.StatusOK, identifier)
	}
}
