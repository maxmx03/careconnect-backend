package medicalprescription

import "database/sql"

type MedicalPrescriptionRepository interface {
	GetAll(doctorID int, patientID int, db *sql.DB) ([]MedicalPrescriptionModel, error)
	Create(medicalPrescription *MedicalPrescriptionModel, db *sql.DB) error
	Update(medicalPrescription *MedicalPrescriptionModel, db *sql.DB) error
	Delete(medicalPrescriptionID int, db *sql.DB) error
}
