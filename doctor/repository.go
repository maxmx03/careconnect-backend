package doctor

import (
	"database/sql"
)

type DoctorRepository interface {
	GetDoctors(db *sql.DB) ([]DoctorModel, error)
	GetDoctorById(doctorID int, db *sql.DB) (*DoctorModel, error)
	CreateDoctor(doctor *DoctorModel, db *sql.DB) error
	UpdateDoctor(doctor *DoctorModel, doctorID int, db *sql.DB) error
	DeleteDoctor(doctorID int, db *sql.DB) error
}
