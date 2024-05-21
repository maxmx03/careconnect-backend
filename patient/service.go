package patient

import (
	"database/sql"
	"errors"
)

type PatientService struct{}

func (s *PatientService) GetAll(db *sql.DB) ([]PatientModel, error) {
	var patients []PatientModel
	query := "SELECT name, surname, cpf, date_of_birth FROM patient"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var patient PatientModel

		if err := rows.Scan(&patient.Name, &patient.Surname, &patient.CPF, &patient.DateOfBirth); err != nil {
			return nil, err
		}

		patients = append(patients, patient)
	}

	return patients, nil
}

func (s *PatientService) GetById(userID int, db *sql.DB) (*PatientModel, error) {
	query := "SELECT name, surname, cpf, date_of_birth FROM patient WHERE user_id = ?"
	patient := &PatientModel{}
	err := db.QueryRow(query, userID).Scan(&patient.Name, &patient.Surname, &patient.CPF, &patient.DateOfBirth)

	if err != nil {
		return patient, err
	}

	return patient, nil
}

func (s *PatientService) Create(patient *PatientModel, db *sql.DB) error {
	query := "INSERT INTO patient (user_id, name, surname, cpf, date_of_birth) VALUES (?, ?, ?, ?, ?)"
	_, err := db.Exec(query, patient.UserID, patient.Name, patient.Surname, patient.CPF, patient.DateOfBirth)

	if err != nil {
		return err
	}

	return nil
}

func (s *PatientService) Update(patient *PatientModel, userID int, db *sql.DB) error {
	query := "UPDATE patient SET name=?, surname=?, cpf=? date_of_birth=? WHERE user_id = ?"
	result, err := db.Exec(query, patient.Name, patient.Surname, patient.CPF, patient.DateOfBirth, userID)

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

func (s *PatientService) Delete(userID int, db *sql.DB) error {
	query := "DELETE FROM patient WHERE patient_id = ?"
	_, err := db.Exec(query, userID)

	if err != nil {
		return err
	}

	return nil
}
