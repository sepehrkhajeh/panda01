package domainapp

import (
	"log"
	"time"

	"github.com/sepehrkhajeh/panda01/infrastructures/database"
	"github.com/sepehrkhajeh/panda01/repositories"

	"github.com/labstack/echo/v4"
)

func RegisterDomainRoutes(e *echo.Echo) {

	cfg := database.Load("config.yaml")

	dbInstance, err := database.Connect(*cfg)
	if err != nil {
		log.Fatal("خطا در اتصال به MongoDB:", err)
	}
	domainRepo := repositories.NewDomainRepository(dbInstance.Client, "domain", 10*time.Second)

	domainGroup := e.Group("/domain")

	domainGroup.GET("/:domain", DetailDomain(domainRepo))
	domainGroup.POST("/create", CreateDomain(domainRepo))
	domainGroup.POST("/:domain", UpdateDomain(domainRepo))
	domainGroup.POST("/delete", DeleteDomain(domainRepo))
}
