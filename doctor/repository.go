package doctor

import (
	"database/sql"
)

type DoctorRepository interface {
	GetDoctors(db *sql.DB) ([]DoctorModel, error)
	GetDoctorById(doctorId int, db *sql.DB) (*DoctorModel, error)
	CreateDoctor(doctor *DoctorModel, db *sql.DB) error
	UpdateDoctor(doctor *DoctorModel, doctorId int, db *sql.DB) error
	DeleteDoctor(doctorId int, db *sql.DB) error
}
