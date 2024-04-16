package user

import (
	"database/sql"
	"errors"
)

type UserService struct{}

func (s *UserService) GetUsers(db *sql.DB) ([]UserModel, error) {
	var users []UserModel
	query := "SELECT * FROM user"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user UserModel
		err := rows.Scan(&user.User_id, &user.Name, &user.Email, &user.Password, &user.Type)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) GetUserById(user *UserModel, db *sql.DB) (*UserModel, error) {
	query := "SELECT * FROM user WHERE user_id = ?"
	err := db.QueryRow(query, user.User_id).Scan(&user.User_id, &user.Name, &user.Password, &user.Email, &user.Type)

	if err != nil {
		return &UserModel{}, err
	}

	return user, nil
}

func (s *UserService) CreateUser(user *UserModel, db *sql.DB) error {
	query := "INSERT INTO user (name, email, password, type) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, user.Name, user.Email, user.Password, user.Type)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(user *UserModel, db *sql.DB) error {
	query := "DELETE FROM user WHERE email = ?"
	_, err := db.Exec(query, user.Email)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUser(user *UserModel, newUser *UserModel, db *sql.DB) error {
	query := "UPDATE user SET name=?, email=?, password=?, type=? WHERE user_id = ?"
	result, err := db.Exec(query, newUser.Name, newUser.Email, newUser.Password, newUser.Type, user.User_id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were updated, user not found")
	}

	return nil
}
