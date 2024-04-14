package user

import (
	"database/sql"
	"errors"
)

type UserService struct{}

func (s *UserService) CreateUser(user *UserModel, db *sql.DB) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	_, err := db.Exec(query, user.Username, user.Email, user.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(email string, db *sql.DB) error {
	query := "DELETE FROM users WHERE email = ?"

	_, err := db.Exec(query, email)

	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserById(id int, db *sql.DB) (UserModel, error) {
	var user UserModel

	query := "SELECT * FROM users WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	if err != nil {
		return UserModel{}, err
	}

	return user, nil
}

func (s *UserService) GetUsers(db *sql.DB) ([]UserModel, error) {
	var users []UserModel

	query := "SELECT * FROM users"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user UserModel

		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) UpdateUser(id int, newUser *UserModel, db *sql.DB) error {

	query := "UPDATE users SET username=?, email=?, password=? WHERE id = ?"

	result, err := db.Exec(query, newUser.Username, newUser.Email, newUser.Password, id)

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
