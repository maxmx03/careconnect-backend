package token

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func ValidateToken(c echo.Context) error {
	token, ok := c.Get("user").(*jwt.Token)

	if ok && token.Valid {
		_, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("failed to cast claims to jwt.MapClaims")
		}
		return nil
	}

	return errors.New("JWT token missing or invalid")
}
