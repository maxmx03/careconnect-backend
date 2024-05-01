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

		if err := rows.Scan(&patient.PatientID, &patient.Cpf, &patient.Username, &patient.Password, &patient.DateOfBirth); err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

func (s *PatientService) GetPatientById(patient *PatientModel, db *sql.DB) (*PatientModel, error) {
	query := "SELECT * FROM patient WHERE patient_id = ?"
	err := db.QueryRow(query, patient.PatientID).Scan(&patient.PatientID, &patient.DateOfBirth, &patient.Username, &patient.Password, &patient.Cpf)

	if err != nil {
		return &PatientModel{}, err
	}

	return patient, nil
}

func (s *PatientService) CreatePatient(patient *PatientModel, db *sql.DB) error {
	query := "INSERT INTO patient (cpf, date_of_birth, username, password) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, patient.Cpf, patient.DateOfBirth, patient.Username, patient.Password)

	if err != nil {
		return err
	}

	return nil
}

func (s *PatientService) DeletePatient(patient *PatientModel, db *sql.DB) error {
	query := "DELETE FROM patient WHERE username = ?"
	_, err := db.Exec(query, patient.Username)

	if err != nil {
		return err
	}

	return nil
}

func (s *PatientService) UpdatePatient(patient *PatientModel, db *sql.DB) error {
	query := "UPDATE patient SET username=?, password=?, date_of_birth=?, cpf=? WHERE patient_id = ?"
	result, err := db.Exec(query, patient.Username, patient.Password, patient.DateOfBirth, patient.Cpf)

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
