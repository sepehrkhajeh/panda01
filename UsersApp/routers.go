package UsersApp

import (
	"Panda/dbconnection"
	"Panda/repositories"
	"log"
	"time"

	"github.com/labstack/echo/v4"
)

// RegisterUserRoutes ثبت روت‌های مربوط به کاربران

func RegisterUserRoutes(e *echo.Echo) {
	client, err := dbconnection.ConnectMongo("mongodb://localhost:27017", 10*time.Second)
	if err != nil {
		log.Fatal("خطا در اتصال به MongoDB:", err)
	}
	userRepo := repositories.NewUserRepository(client, "users", 10*time.Second)
	e.POST("/users", CreateUser(userRepo))
	e.POST("/deleteuser", DeleteUser(userRepo))
}
