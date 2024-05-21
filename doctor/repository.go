package doctor

import (
	"database/sql"
)

type DoctorRepository interface {
	GetAll(db *sql.DB) ([]DoctorModel, error)
	GetById(userID int, db *sql.DB) (*DoctorModel, error)
	Create(doctor *DoctorModel, db *sql.DB) error
	Update(doctor *DoctorModel, userID int, db *sql.DB) error
	Delete(userID int, db *sql.DB) error
}
