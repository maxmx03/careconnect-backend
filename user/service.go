package user

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/maxmx03/careconnect-backend/token"
)

type UserService struct{}

func (s *UserService) Create(user *UserModel, db *sql.DB) error {
	query := "INSERT INTO user (username, password, user_type) VALUES (?, ?, ?)"
	_, err := db.Exec(query, user.Username, user.Password, user.UserType)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Delete(userID int, db *sql.DB) error {
	query := "DELETE FROM user where user_id = ?"
	_, err := db.Exec(query, userID)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) Update(user *UserModel, userID int, db *sql.DB) error {
	query := "UPDATE user SET username=?, password=?, user_type=? WHERE user_id = ?"
	result, err := db.Exec(query, user.Username, user.Password, userID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No rows were updated, user not found")
	}

	return nil
}

func (s *UserService) Login(auth *UserModel, db *sql.DB) (string, error) {
	user := &UserModel{}
	var query string

	query = "SELECT username, password FROM user WHERE username = ?"

	if err := db.QueryRow(query, auth.Username).Scan(&user.Username, &user.Password); err != nil {
		return "", err
	}

	if auth.Username == user.Username && auth.Password == user.Password {
		t, err := token.Create(user.UserID, user.UserType)

		if err != nil {
			return "", err
		}

		return t, nil
	}

	return "", errors.New("Unauthorized")
}
