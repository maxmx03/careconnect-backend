package medicalprescription

import "database/sql"

type MedicalPrescriptionRepository interface {
	GetMedicalPrescriptions(doctorID int, patientID int, db *sql.DB) ([]MedicalPrescriptionModel, error)
	CreateMedicalPrescription(medicalPrescription *MedicalPrescriptionModel, db *sql.DB) error
	UpdateMedicalPrescription(medicalPrescription *MedicalPrescriptionModel, db *sql.DB) error
	DeleteMedicalPrescription(medicalPrescriptionID int, db *sql.DB) error
}
