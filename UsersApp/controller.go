package usersapp

import (
	"net/http"

	"github.com/sepehrkhajeh/panda01/repositories"
	"github.com/sepehrkhajeh/panda01/schemas"
	"github.com/sepehrkhajeh/panda01/validations"

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

		_, err = repo.Add(ctx, &schemas.UserSchema{PubKey: user.PubKey})
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
		if _, err := repo.Delete(ctx, d); err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Error deleting user")
		}

		return c.JSON(http.StatusGone, map[string]string{"message": "User successfully deleted"})
	}
}
