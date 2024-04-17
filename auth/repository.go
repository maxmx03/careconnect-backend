package auth

import (
	"database/sql"
	. "github.com/maxmx03/careconnect-backend/user"
)

type AuthRepository interface {
	Login(user *UserModel, db *sql.DB) (string, error)
}
