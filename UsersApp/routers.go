package usersapp

import (
	"log"
	"time"

	"github.com/sepehrkhajeh/panda01/dbconnection"
	"github.com/sepehrkhajeh/panda01/repositories"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	client, err := dbconnection.ConnectMongo("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal("خطا در اتصال به MongoDB:", err)
	}
	userRepo := repositories.NewUserRepository(client, "users", 10*time.Second)
	e.POST("/users", CreateUser(userRepo))
	e.POST("/deleteuser", DeleteUser(userRepo))
}
