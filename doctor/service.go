package doctor

import (
	"database/sql"
	"errors"
)

type DoctorService struct{}

func (s *DoctorService) GetDoctors(db *sql.DB) ([]DoctorModel, error) {
	var doctors []DoctorModel
	query := "SELECT doctor_id, crm, username, password FROM doctor"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var doctor DoctorModel

		if err := rows.Scan(&doctor.Doctor_id, &doctor.Crm, &doctor.Username, &doctor.Password); err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (s *DoctorService) GetDoctorById(doctor *DoctorModel, db *sql.DB) (*DoctorModel, error) {
	query := "SELECT doctor_id, crm, username, password  FROM doctor WHERE doctor_id = ?"
	err := db.QueryRow(query, doctor.Doctor_id).Scan(&doctor.Doctor_id, &doctor.Crm, &doctor.Username, &doctor.Password)

	if err != nil {
		return &DoctorModel{}, err
	}

	return doctor, nil
}

func (s *DoctorService) CreateDoctor(doctor *DoctorModel, db *sql.DB) error {
	query := "INSERT INTO doctor (username, password, crm) VALUES (?, ?, ?)"
	_, err := db.Exec(query, doctor.Username, doctor.Password, doctor.Crm)

	if err != nil {
		return err
	}

	return nil
}

func (s *DoctorService) DeleteDoctor(doctor *DoctorModel, db *sql.DB) error {
	query := "DELETE FROM doctor WHERE username = ?"
	_, err := db.Exec(query, doctor.Username)

	if err != nil {
		return err
	}

	return nil
}

func (s *DoctorService) UpdateDoctor(doctor *DoctorModel, db *sql.DB) error {
	query := "UPDATE user SET username=?, password=?, crm=? WHERE doctor_id = ?"
	result, err := db.Exec(query, doctor.Username, doctor.Password, doctor.Password, doctor.Crm, doctor.Doctor_id)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were updated, user not found")
	}

	return nil
}
