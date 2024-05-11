package doctor

import (
	"database/sql"
	"errors"
)

type DoctorService struct{}

func (s *DoctorService) GetDoctors(db *sql.DB) ([]DoctorModel, error) {
	var doctors []DoctorModel
	query := "SELECT doctor_id, name, crm, username, password FROM doctor"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var doctor DoctorModel

		if err := rows.Scan(&doctor.DoctorID, &doctor.Name, &doctor.CRM, &doctor.Username, &doctor.Password); err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (s *DoctorService) GetDoctorById(doctorId int, db *sql.DB) (*DoctorModel, error) {
	query := "SELECT doctor_id, name, crm, username, password  FROM doctor WHERE doctor_id = ?"
	doctor := &DoctorModel{}
	err := db.QueryRow(query, doctorId).Scan(&doctor.DoctorID, &doctor.Name, &doctor.CRM, &doctor.Username, &doctor.Password)

	if err != nil {
		return doctor, err
	}

	return doctor, nil
}

func (s *DoctorService) CreateDoctor(doctor *DoctorModel, db *sql.DB) error {
	query := "INSERT INTO doctor (name, username, password, crm) VALUES (?, ?, ?)"
	_, err := db.Exec(query, doctor.Name, doctor.Username, doctor.Password, doctor.CRM)

	if err != nil {
		return err
	}

	return nil
}

func (s *DoctorService) DeleteDoctor(doctorId int, db *sql.DB) error {
	query := "DELETE FROM doctor WHERE doctor_id = ?"
	_, err := db.Exec(query, doctorId)

	if err != nil {
		return err
	}

	return nil
}

func (s *DoctorService) UpdateDoctor(doctor *DoctorModel, doctorId int, db *sql.DB) error {
	query := "UPDATE doctor SET name=? username=?, password=?, crm=? WHERE doctor_id = ?"
	result, err := db.Exec(query, doctor.Name, doctor.Username, doctor.Password, doctor.CRM, doctorId)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no rows were updated, doctor not found")
	}

	return nil
}
