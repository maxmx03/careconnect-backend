package user

import (
	"database/sql"
)

type UserRepository interface {
	GetUsers(db *sql.DB) ([]UserModel, error)
	GetUserById(id int, db *sql.DB) (UserModel, error)
	CreateUser(user *UserModel, db *sql.DB) error
	UpdateUser(id int, newUser *UserModel, db *sql.DB) error
	DeleteUser(email string, db *sql.DB) error
}
