package main

import (
	"Panda/Users"
	"Panda/dbconnection"
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	collection, err := dbconnection.Collection()
	if err != nil {
		return
	}

	e := echo.New()
	e.GET("/.well-known/nostr.json", func(c echo.Context) error {
		name := c.QueryParam("name")
		if name == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "name parameter is required"})
		}
		var user Users.User
		err := collection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return c.JSON(http.StatusNotFound, map[string]string{"error": "user not found"})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}

		response := map[string]interface{}{
			"names": map[string]string{
				user.Name: user.PubKey,
			},
			"relays": map[string][]string{
				user.PubKey: user.Relays,
			},
		}

		return c.JSON(http.StatusOK, response)
	})

	e.POST("/users", func(c echo.Context) error {

		var user Users.UserJs
		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
		}

		// افزودن سند به کالکشن
		_, err := collection.InsertOne(context.TODO(), user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not insert user"})
		}

		return c.JSON(http.StatusCreated, user)
	})
	e.Logger.Fatal(e.Start(":8000"))
}
