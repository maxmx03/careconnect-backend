package auth

import (
	"database/sql"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	. "github.com/maxmx03/careconnect-backend/user"
)

type jwtCustomClaims struct {
	Email string `json:"Email"`
	Type  string
	jwt.RegisteredClaims
}

type AuthService struct{}

func (s *AuthService) Login(user *UserModel, db *sql.DB) (string, error) {
	query := "SELECT email, password, type FROM user where user_id = ?"
	dbUser := &UserModel{}

	if err := db.QueryRow(query, user.User_id).Scan(&dbUser.Email, &dbUser.Password, &dbUser.Type); err != nil {
		return "", err
	}

	if user.Email == dbUser.Email && user.Password == dbUser.Password {
		claims := &jwtCustomClaims{
			dbUser.Email,
			dbUser.Type,
			jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

		t, err := token.SignedString([]byte("secret"))

		if err != nil {
			return "", err
		}

		return t, nil
	}

	return "", errors.New("unauthorized")
}
