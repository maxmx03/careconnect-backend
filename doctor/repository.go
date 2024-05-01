package doctor

import (
	"database/sql"
)

type DoctorRepository interface {
	GetDoctors(db *sql.DB) ([]DoctorModel, error)
	GetDoctorById(doctor *DoctorModel, db *sql.DB) (*DoctorModel, error)
	CreateDoctor(doctor *DoctorModel, db *sql.DB) error
	UpdateDoctor(doctor *DoctorModel, db *sql.DB) error
	DeleteDoctor(doctor *DoctorModel, db *sql.DB) error
}
