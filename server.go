package main

import (
	"Panda/DomainApp"
	"Panda/IdentifierApp"
	"Panda/UsersApp"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	UsersApp.RegisterUserRoutes(e)
	DomainApp.RegisterDomainRoutes(e)
	IdentifierApp.RegisterIdentifierRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
