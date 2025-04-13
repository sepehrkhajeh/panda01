package usersapp

import (
	"log"
	"time"

	"github.com/sepehrkhajeh/panda01/infrastructures/database"
	"github.com/sepehrkhajeh/panda01/repositories"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {

	cfg := database.Load("config.yaml")

	dbInstance, err := database.Connect(*cfg)
	if err != nil {
		log.Fatal("خطا در اتصال به MongoDB:", err)
	}
	userRepo := repositories.NewUserRepository(dbInstance.Client, "users", 10*time.Second)

	e.POST("/users", CreateUser(userRepo))
	e.POST("/deleteuser", DeleteUser(userRepo))
}
