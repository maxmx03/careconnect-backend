package medicalconsultation

import (
	"database/sql"
)

type MedicalConsultationRepository interface {
	GetAll(doctorID int, patientID int, db *sql.DB) ([]MedicalConsultationModel, error)
	Create(consultation *MedicalConsultationModel, db *sql.DB) error
	Update(consultation *MedicalConsultationModel, db *sql.DB) error
	Delete(consultationID int, db *sql.DB) error
}
