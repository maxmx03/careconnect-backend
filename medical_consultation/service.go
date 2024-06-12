package medicalconsultation

import (
	"database/sql"
	"errors"
)

type MedicalConsultationService struct{}

func (r *MedicalConsultationService) GetAll(doctorID int, patientID int, db *sql.DB) ([]MedicalConsultationModel, error) {
	var medicalConsultations []MedicalConsultationModel
	query := `
    SELECT
        mc.consultation_id,
        mc.doctor_id,
        mc.patient_id,
        mc.date_time,
        mc.description,
        d.name AS doctor_name,
        d.surname AS doctor_surname,
        p.name AS patient_name,
        p.surname AS patient_surname
    FROM
        medical_consultation mc
    INNER JOIN
        doctor d ON mc.doctor_id = d.doctor_id
    INNER JOIN
        patient p ON mc.patient_id = p.patient_id
    WHERE
        d.doctor_id = ? AND p.patient_id = ?
    `
	rows, err := db.Query(query, doctorID, patientID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var medicalConsultation MedicalConsultationModel

		if err := rows.Scan(
			&medicalConsultation.ConsultationID,
			&medicalConsultation.DoctorID,
			&medicalConsultation.PatientID,
			&medicalConsultation.DateTime,
			&medicalConsultation.Description,
			&medicalConsultation.DoctorName,
			&medicalConsultation.DoctorSurname,
			&medicalConsultation.PatientName,
			&medicalConsultation.PatientSurname,
		); err != nil {
		}
		return nil, err
	}

	return medicalConsultations, nil
}

func (r *MedicalConsultationService) Create(consultation *MedicalConsultationModel, db *sql.DB) error {
	query := `
    INSERT INTO medical_consultation (doctor_id, patient_id, date_time, description)
    VALUES (?, ?, ?, ?);
    `
	result, err := db.Exec(query, consultation.DoctorID, consultation.PatientID, consultation.DateTime, consultation.Description)
	if err != nil {
		return err
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}
	consultation.ConsultationID = int(lastInsertID)
	return nil
}

func (r *MedicalConsultationService) Update(consultation *MedicalConsultationModel, db *sql.DB) error {
	query := `
    UPDATE medical_consultation
    SET doctor_id = ?, patient_id = ?, date_time = ?, description = ?
    WHERE consultation_id = ?;
    `
	result, err := db.Exec(query, consultation.DoctorID, consultation.PatientID, consultation.DateTime, consultation.Description, consultation.ConsultationID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were updated, consultation not found")
	}

	return nil
}

func (r MedicalConsultationService) Delete(consultationID int, db *sql.DB) error {
	query := "DELETE FROM medical_consultation WHERE consultation_id = ?"
	_, err := db.Exec(query, consultationID)

	if err != nil {
		return err
	}

	return nil
}
