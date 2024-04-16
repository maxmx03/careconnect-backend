package doctor

import (
	"database/sql"

	. "github.com/maxmx03/careconnect-backend/user"
)

type DoctorRepository interface {
  GetDoctor(user *UserModel, db *sql.DB) (*DoctorModel, error)
}
