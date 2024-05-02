package auth

import (
	"database/sql"
	"errors"
	"github.com/maxmx03/careconnect-backend/token"
)

type AuthService struct{}

func (s *AuthService) Login(auth *AuthModel, db *sql.DB) (string, error) {
	user := &AuthModel{}
	var query string

	if auth.Type == "doctor" {
		query = "SELECT username, password FROM doctor where username = ?"
	} else {
		query = "SELECT username, password FROM patient where username = ?"
	}

	if err := db.QueryRow(query, auth.Username).Scan(&user.Username, &user.Password); err != nil {
		return "", err
	}

	if auth.Username == user.Username && auth.Password == user.Password {
		t, err := token.CreateToken(user.Username, user.Type)

		if err != nil {
			return "", err
		}

		return t, nil
	}

	return "", errors.New("Unauthorized")
}
