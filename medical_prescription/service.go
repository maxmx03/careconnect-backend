package medicalprescription

import "database/sql"

type MedicalPrescriptionService struct{}

func (s *MedicalPrescriptionService) GetMedicalPrescriptionsById(doctorID int, patientID int, db *sql.DB) ([]MedicalPrescriptionModel, error) {
	var medicalPrescriptions []MedicalPrescriptionModel
	query := `
SELECT
    medical_prescription.prescription_id, 
    medical_prescription.date, 
    medical_prescription.description,
    d.doctor_id, 
    d.name AS doctor_name, 
    d.surname AS doctor_surname, 
    p.patient_id, 
    p.name AS patient_name, 
    p.surname AS patient_surname
FROM 
    medical_prescription
INNER JOIN 
    doctor d ON medical_prescription.doctor_id = d.id
INNER JOIN 
    patient p ON medical_prescription.patient_id = p.patient_id
WHERE 
    d.doctor_id = ? AND p.patient_id = ?
  `
	rows, err := db.Query(query, doctorID, patientID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var medicalPrescription MedicalPrescriptionModel

		if err := rows.Scan(&medicalPrescription.DoctorID, &medicalPrescription.Date, &medicalPrescription.Description, &medicalPrescription.DoctorID,
			&medicalPrescription.DoctorName, &medicalPrescription.DoctorSurname, &medicalPrescription.PatientID,
			&medicalPrescription.PatientName, &medicalPrescription.PatientSurname); err != nil {
			return nil, err
		}

		medicalPrescriptions = append(medicalPrescriptions, medicalPrescription)
	}

	return medicalPrescriptions, nil
}

func (s *MedicalPrescriptionService) CreateMedicalPrescription(medicalPrescription *MedicalPrescriptionModel, db *sql.DB) error {
	query := `
    INSERT INTO medical_prescription (date, description, doctor_id, patient_id)
    VALUES (?, ?, ?, ?)
    `
	_, err := db.Exec(query, medicalPrescription.Date, medicalPrescription.Description, medicalPrescription.DoctorID, medicalPrescription.PatientID)
	if err != nil {
		return err
	}
	return nil
}

func (s *MedicalPrescriptionService) UpdateMedicalPrescription(medicalPrescription *MedicalPrescriptionModel, db *sql.DB) error {
	query := `
    UPDATE medical_prescription
    SET date = ?, description = ?
    WHERE prescription_id = ?
    `
	_, err := db.Exec(query, medicalPrescription.Date, medicalPrescription.Description, medicalPrescription.PrescriptionID)
	if err != nil {
		return err
	}
	return nil
}

func (s *MedicalPrescriptionService) DeleteMedicalPrescription(prescriptionID int, db *sql.DB) error {
	query := `
    DELETE FROM medical_prescription
    WHERE prescription_id = ?
    `
	_, err := db.Exec(query, prescriptionID)
	if err != nil {
		return err
	}
	return nil
}
