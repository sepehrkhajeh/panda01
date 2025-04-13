package domainapp

import (
	"log"
	"net/http"
	"time"

	"github.com/sepehrkhajeh/panda01/repositories"
	"github.com/sepehrkhajeh/panda01/schemas"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
)

func CreateDomain(repo *repositories.DomainRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(DomainCreateRequest)
		if err := c.Bind(req); err != nil {
			return err
		}
		ctx := c.Request().Context()
		validationErrors := ValidateData(*req)
		if validationErrors != nil {
			return c.JSON(http.StatusBadRequest, validationErrors)
		}
		d, err := repo.GetByFeild(ctx, "domain", req.Domain)
		if err != nil {
			return err
		}
		if d != nil {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": "Domain already exists",
			})
		}

		_, err = repo.Add(ctx, &schemas.DomainSchema{
			Domain:                 req.Domain,
			BasePricePerIdentifier: req.BasePricePerIdentifier,
			DefaultTTL:             req.DefaultTTL,
			Status:                 req.Status,
		})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusCreated, req)

	}
}

func DetailDomain(repo *repositories.DomainRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		domain := c.Param("domain")
		ctx := c.Request().Context()
		d, err := repo.GetByFeild(ctx, "domain", domain)
		if d == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "not found domain"})
		}
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, d)
	}
}

func UpdateDomain(repo *repositories.DomainRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		updateData := make(map[string]interface{})
		if err := c.Bind(&updateData); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		ctx := c.Request().Context()
		domainParam := c.Param("domain")
		updateData["updated_at"] = time.Now()
		d, err := repo.GetByFeild(ctx, "domain", domainParam)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if d == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Domain does not exist"})
		}
		update := bson.M{"$set": updateData}

		_, err = repo.Update(ctx, bson.M{"domain": domainParam}, update)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, updateData)
	}
}

func DeleteDomain(repo *repositories.DomainRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		domain := make(map[string]interface{})
		if err := c.Bind(&domain); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		if domain["domain"] == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "domain is required"})
		}
		log.Printf("domain : %v", domain["domain"])

		ctx := c.Request().Context()
		d, err := repo.GetByFeild(ctx, "domain", domain["domain"])
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		if d == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Domain does not exist"})
		}
		result, err := repo.Delete(ctx, bson.M{"domain": domain["domain"]})
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		log.Printf("reslut ==========> : %v", result)
		return c.JSON(http.StatusOK, map[string]string{"message": "domain has been deleted."})
	}
}
