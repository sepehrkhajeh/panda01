package identifierapp

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sepehrkhajeh/panda01/repositories"
	"github.com/sepehrkhajeh/panda01/schemas"
)

func CreateIdentifier(identifierRepo *repositories.IdentifierRepository, domainRepo *repositories.DomainRepository) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := new(IdentifierCreateRequest)

		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
		}
		ctx := c.Request().Context()

		domainData, err := domainRepo.GetByID(ctx, req.DomainID)
		log.Print(domainData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error retrieving domain information"})
		}
		if domainData == nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "domain not found"})
		}

		fullIdentifier := fmt.Sprintf("%s@%s", req.Name, domainData.Domain)

		log.Println(fullIdentifier)
		newIdentifier := schemas.IdentifierSchema{
			Name:           req.Name,
			Pubkey:         req.Pubkey,
			DomainID:       req.DomainID,
			FullIdentifier: fullIdentifier,
		}

		reslut, err := identifierRepo.IsExist(ctx, "full_identifier", fullIdentifier)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error retrieving identifier information"})
		}
		if reslut {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "identifier already exists"})
		}
		identifier, err := identifierRepo.Add(ctx, newIdentifier)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "error creating identifier"})
		}

		return c.JSON(http.StatusOK, identifier)
	}
}
