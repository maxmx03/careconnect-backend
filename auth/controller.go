package auth

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	. "github.com/maxmx03/careconnect-backend/message"
)

type AuthController struct{}

var authService AuthRepository = &AuthService{}

func (a *AuthController) Login(c echo.Context, db *sql.DB) error {
	auth := &AuthModel{}
	var err error
	var token string

	if err = c.Bind(auth); err != nil {
		return err
	}

	if token, err = authService.Login(auth, db); err != nil {
		return c.JSON(http.StatusForbidden, GetError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}
