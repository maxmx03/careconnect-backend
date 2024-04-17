package auth

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var authController = &AuthController{}

func AuthRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/login", func(c echo.Context) error {
		return authController.Login(c, db)
	})
}
