package patient

import (
	"database/sql"
	"errors"
)

type PatientService struct{}

func (s *PatientService) GetPatients(db *sql.DB) ([]PatientModel, error) {
	var patients []PatientModel
	query := "SELECT * FROM patient"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient PatientModel

		if err := rows.Scan(&patient.PatientID, &patient.Name, &patient.CPF, &patient.Username, &patient.Password, &patient.DateOfBirth); err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

func (s *PatientService) GetPatientById(patientID int, db *sql.DB) (*PatientModel, error) {
	query := "SELECT * FROM patient WHERE patient_id = ?"
	patient := &PatientModel{}
	err := db.QueryRow(query, patientID).Scan(&patient.PatientID, &patient.Name, &patient.DateOfBirth, &patient.Username, &patient.Password, &patient.CPF)

	if err != nil {
		return patient, err
	}

	return patient, nil
}

func (s *PatientService) CreatePatient(patient *PatientModel, db *sql.DB) error {
	query := "INSERT INTO patient (name, cpf, date_of_birth, username, password) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, patient.Name, patient.CPF, patient.DateOfBirth, patient.Username, patient.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *PatientService) DeletePatient(patientID int, db *sql.DB) error {
	query := "DELETE FROM patient WHERE patient_id = ?"
	_, err := db.Exec(query, patientID)

	if err != nil {
		return err
	}

	return nil
}

func (s *PatientService) UpdatePatient(patient *PatientModel, db *sql.DB) error {
	query := "UPDATE patient SET name=?, username=?, password=?, date_of_birth=?, cpf=? WHERE patient_id = ?"
	result, err := db.Exec(query, patient.Name, patient.Username, patient.Password, patient.DateOfBirth, patient.CPF)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were updated, patient not found")
	}

	return nil
}
