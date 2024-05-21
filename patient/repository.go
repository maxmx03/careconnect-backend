package patient

import (
	"database/sql"
)

type PatientRepository interface {
	GetAll(db *sql.DB) ([]PatientModel, error)
	GetById(userID int, db *sql.DB) (*PatientModel, error)
	Create(patient *PatientModel, db *sql.DB) error
	Update(patient *PatientModel, userID int, db *sql.DB) error
	Delete(userID int, db *sql.DB) error
}
