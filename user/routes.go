package user

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var userController = &UserController{}

func Routes(e *echo.Echo, db *sql.DB) {
	e.GET("/users", func(c echo.Context) error {
		return userController.GetUsers(c, db)
	})
	e.GET("/user", func(c echo.Context) error {
		return userController.GetUserById(c, db)
	})
	e.POST("/user", func(c echo.Context) error {
		return userController.CreateUser(c, db)
	})
	e.PUT("/user", func(c echo.Context) error {
		return userController.UpdateUser(c, db)
	})
	e.DELETE("/user", func(c echo.Context) error {
		return userController.DeleteUser(c, db)
	})
}
