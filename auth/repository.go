package auth

import (
	"database/sql"
)

type AuthRepository interface {
	Login(user *AuthModel, db *sql.DB) (string, error)
}
