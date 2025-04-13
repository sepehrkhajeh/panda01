package main

import (
	domainapp "github.com/sepehrkhajeh/panda01/app/domain"
	identifierapp "github.com/sepehrkhajeh/panda01/app/identifier"
	usersapp "github.com/sepehrkhajeh/panda01/app/users"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	usersapp.RegisterUserRoutes(e)
	domainapp.RegisterDomainRoutes(e)
	identifierapp.RegisterIdentifierRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}
