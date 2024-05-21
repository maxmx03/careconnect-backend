package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

type JwtCustomClaims struct {
	UserID   int `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.RegisteredClaims
}

type SignKey string

func validate(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_, ok := user.Claims.(*JwtCustomClaims)

	if !ok {
		return errors.New("Invalid token")
	}

	return nil
}

func Create(userID int, userType string) (string, error) {
	privateKey, err := os.ReadFile("ec256-private.pem")

	if err != nil {
		panic(err)
	}

	claims := &JwtCustomClaims{
		userID,
		userType,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return t, nil
}

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}

		if err := validate(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
