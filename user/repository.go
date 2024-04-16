package user

import (
	"database/sql"
)

type UserRepository interface {
	GetUsers(db *sql.DB) ([]UserModel, error)
	GetUserById(user *UserModel, db *sql.DB) (*UserModel, error)
	CreateUser(user *UserModel, db *sql.DB) error
	UpdateUser(user *UserModel, newUser *UserModel, db *sql.DB) error
	DeleteUser(user *UserModel, db *sql.DB) error
}
