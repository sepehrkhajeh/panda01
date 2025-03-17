package identifierapp

import (
	"log"
	"time"

	"github.com/sepehrkhajeh/panda01/dbconnection"
	"github.com/sepehrkhajeh/panda01/repositories"

	"github.com/labstack/echo/v4"
)

func RegisterIdentifierRoutes(e *echo.Echo) {
	client, err := dbconnection.ConnectMongo("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal("خطا در اتصال به MongoDB:", err)
	}

	IdentifierRepo := repositories.NewIdentifierRepository(client, "identifier", 10*time.Second)

	identifierGroup := e.Group("/indentifier")

	identifierGroup.POST("/add", CreateIdentifier(IdentifierRepo))

}
