package user

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var userController = &UserController{}

func Routes(e *echo.Echo, db *sql.DB) {
	var url = "/users"
	g := e.Group(url)
	g.POST("/login", func(c echo.Context) error {
		return userController.Login(c, db)
	})
	g.POST("", func(c echo.Context) error {
		return userController.Create(c, db)
	})
	g.PUT("/:id", func(c echo.Context) error {
		return userController.Update(c, db)
	})
	g.DELETE("/:id", func(c echo.Context) error {
		return userController.Delete(c, db)
	})
}
