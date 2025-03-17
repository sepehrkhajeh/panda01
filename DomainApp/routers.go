package domainapp

import (
	"log"
	"time"

	"github.com/sepehrkhajeh/panda01/dbconnection"
	"github.com/sepehrkhajeh/panda01/repositories"

	"github.com/labstack/echo/v4"
)

func RegisterDomainRoutes(e *echo.Echo) {
	client, err := dbconnection.ConnectMongo("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal("خطا در اتصال به MongoDB:", err)
	}

	domainRepo := repositories.NewDomainRepository(client, "domain", 10*time.Second)

	domainGroup := e.Group("/domain")

	domainGroup.GET("/:domain", DetailDomain(domainRepo))
	domainGroup.POST("/create", CreateDomain(domainRepo))
	domainGroup.POST("/:domain", UpdateDomain(domainRepo))
	domainGroup.POST("/delete", DeleteDomain(domainRepo))
}
