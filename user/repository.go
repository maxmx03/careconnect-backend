package user

import (
	"database/sql"
)

type UserRepository interface {
	Login(user *UserModel, db *sql.DB) (string, error)
	Create(user *UserModel, db *sql.DB) error
	Update(user *UserModel, userID int, db *sql.DB) error
	Delete(userID int, db *sql.DB) error
}
