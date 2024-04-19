package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maxmx03/careconnect-backend/token"
)

var userController = &UserController{}

func UserRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/users", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return userController.GetUsers(c, db)
	})
	e.GET("/user", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return userController.GetUserById(c, db)
	})
	e.POST("/user", func(c echo.Context) error {
		return userController.CreateUser(c, db)
	})
	e.PUT("/user", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return userController.UpdateUser(c, db)
	})
	e.DELETE("/user", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return userController.DeleteUser(c, db)
	})
}
