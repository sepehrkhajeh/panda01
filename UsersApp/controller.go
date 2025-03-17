package UsersApp

import (
	"Panda/repositories"
	"Panda/schemas"
	"Panda/validations"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateUser(repo *repositories.UserRepasitory) echo.HandlerFunc {

	return func(c echo.Context) error {

		user := new(schemas.UserSchema)
		if err := c.Bind(user); err != nil {
			return err
		}
		
		validate := validations.NewValidator()
		if err := validate.Struct(user); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		ctx := c.Request().Context()

		d, err := repo.GetByFeild(ctx, "pubKey", user.PubKey)
		if err != nil {
			return err
		}
		if d != nil {
			return c.JSON(http.StatusConflict, map[string]string{
				"error": "user already exists",
			})
		}

		log.Printf("داده ورودی: %+v", user)
		log.Printf("داده ورودی: %+v", c.Request().Body)
		log.Printf("ctx {}--- %+v", ctx)
		_, err = repo.Add(ctx, &schemas.UserSchema{PubKey: user.PubKey})
		log.Printf("===================================%+v", ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		return c.JSON(http.StatusOK, user)
	}
}

func DeleteUser(repo *repositories.UserRepasitory) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := new(schemas.UserSchema)
		if err := c.Bind(user); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid input")
		}
		ctx := c.Request().Context()
		d, err := repo.GetByFeild(ctx, "pubKey", user.PubKey)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error fetching user")
		}
		if d == nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "User does not exist",
			})
		}
		log.Printf("===================================%+v", d)
		if _, err := repo.Delete(ctx, d); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting user")
		}

		return c.JSON(http.StatusGone, map[string]string{"message": "User successfully deleted"})
	}
}
