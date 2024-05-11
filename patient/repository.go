package patient

import (
	"database/sql"
)

type PatientRepository interface {
	GetPatients(db *sql.DB) ([]PatientModel, error)
	GetPatientById(patientID int, db *sql.DB) (*PatientModel, error)
	CreatePatient(patient *PatientModel, db *sql.DB) error
	UpdatePatient(patient *PatientModel, db *sql.DB) error
	DeletePatient(patientID int, db *sql.DB) error
}
