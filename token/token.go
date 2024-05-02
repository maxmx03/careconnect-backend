package token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"os"
	"time"
)

type JwtCustomClaims struct {
	Name     string `json:"name"`
	UserType string `json:"admin"`
	jwt.RegisteredClaims
}

type SignKey string

func ValidateToken(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	_, ok := user.Claims.(*JwtCustomClaims)

	if !ok {
		return errors.New("Invalid token")
	}

	return nil
}

func CreateToken(email string, userType string) (string, error) {
	privateKey, err := os.ReadFile("ec256-private.pem")

	if err != nil {
		panic(err)
	}

	claims := &JwtCustomClaims{
		email,
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
