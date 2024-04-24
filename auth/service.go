package auth

import (
	"database/sql"
	"errors"
	"github.com/maxmx03/careconnect-backend/token"
	. "github.com/maxmx03/careconnect-backend/user"
)

type AuthService struct{}

func (s *AuthService) Login(auth *AuthModel, db *sql.DB) (string, error) {
	query := "SELECT email, password, type FROM user where email = ?"
	user := &UserModel{}

	if err := db.QueryRow(query, auth.Email).Scan(&user.Email, &user.Password, &user.Type); err != nil {
		return "", err
	}

	if auth.Email == user.Email && auth.Password == user.Password {
		t, err := token.CreateToken(user.Email, user.Type)

		if err != nil {
			return "", err
		}

		return t, nil
	}

	return "", errors.New("Unauthorized")
}
