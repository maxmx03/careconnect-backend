package auth

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	. "github.com/maxmx03/careconnect-backend/user"
)

type AuthController struct{}

var authService AuthRepository = &AuthService{}

func (a *AuthController) Login(c echo.Context, db *sql.DB) error {
	user := &UserModel{}
	var err error
	var token string

	if err = c.Bind(user); err != nil {
		return err
	}

	if token, err = authService.Login(user, db); err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

