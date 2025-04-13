package identifierapp

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"

	"github.com/sepehrkhajeh/panda01/infrastructures/database"
	"github.com/sepehrkhajeh/panda01/repositories"
)

func RegisterIdentifierRoutes(e *echo.Echo) {

	cfg := database.Load("config.yaml")
	dbInstance, err := database.Connect(*cfg)
	if err != nil {
		log.Fatal("خطا در اتصال به MongoDB:", err)
	}
	identifierRepo := repositories.NewIdentifierRepository(dbInstance.Client, "identifier", 10*time.Second)
	domainRepo := repositories.NewDomainRepository(dbInstance.Client, "domain", 10*time.Second)

	identifierGroup := e.Group("/identifier")
	identifierGroup.POST("/add", CreateIdentifier(identifierRepo, domainRepo))
}
