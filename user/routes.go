package user

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var userController = &UserController{}

func Routes(e *echo.Echo, db *sql.DB) {
	e.POST("/login", func(c echo.Context) error {
		return userController.Login(c, db)
	})
}
