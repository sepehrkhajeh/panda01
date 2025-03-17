package main

import (
	"github.com/sepehrkhajeh/panda01/domainapp"
	"github.com/sepehrkhajeh/panda01/identifierapp"
	"github.com/sepehrkhajeh/panda01/usersapp"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	usersapp.RegisterUserRoutes(e)
	domainapp.RegisterDomainRoutes(e)
	identifierapp.RegisterIdentifierRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
