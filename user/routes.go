package user

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo, db *sql.DB) {
	UserController := &UserController{}

	e.GET("/users", func(c echo.Context) error {
		return UserController.GetUsers(c, db)
	})
	e.GET("/user", func(c echo.Context) error {
		return UserController.GetUserById(c, db)
	})
	e.POST("/user", func(c echo.Context) error {
		return UserController.CreateUser(c, db)
	})
	e.PUT("/user", func(c echo.Context) error {
		return UserController.UpdateUser(c, db)
	})
	e.DELETE("/user", func(c echo.Context) error {
		return UserController.DeleteUser(c, db)
	})
}
