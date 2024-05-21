package user

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	. "github.com/maxmx03/careconnect-backend/feedback"
)

type UserController struct{}

var userService UserRepository = &UserService{}

func (a *UserController) Login(c echo.Context, db *sql.DB) error {
	auth := &UserModel{}
	var err error
	var token string

	if err = c.Bind(auth); err != nil {
		return err
	}

	if token, err = userService.Login(auth, db); err != nil {
		return c.JSON(http.StatusForbidden, GetError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
