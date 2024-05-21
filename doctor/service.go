package doctor

import (
	"database/sql"
	"errors"
)

type DoctorService struct{}

func (s *DoctorService) GetAll(db *sql.DB) ([]DoctorModel, error) {
	var doctors []DoctorModel
	query := "SELECT name, surname, crm FROM doctor"
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var doctor DoctorModel

		if err := rows.Scan(&doctor.Name, &doctor.Surname, &doctor.CRM); err != nil {
			return nil, err
		}

		doctors = append(doctors, doctor)
	}

	return doctors, nil
}

func (s *DoctorService) GetById(userID int, db *sql.DB) (*DoctorModel, error) {
	query := "SELECT name, surname, crm FROM doctor WHERE user_id = ?"
	doctor := &DoctorModel{}
	err := db.QueryRow(query, userID).Scan(&doctor.Name, &doctor.Surname, &doctor.CRM)

	if err != nil {
		return doctor, err
	}

	return doctor, nil
}

func (s *DoctorService) Create(doctor *DoctorModel, db *sql.DB) error {
	query := "INSERT INTO doctor (user_id, name, surname, crm) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(query, doctor.UserID, doctor.Name, doctor.Surname, doctor.CRM)

	if err != nil {
		return err
	}

	return nil
}

func (s *DoctorService) Update(doctor *DoctorModel, userID int, db *sql.DB) error {
	query := "UPDATE doctor SET name=?, surname=?, crm=? WHERE user_id = ?"
	result, err := db.Exec(query, doctor.Name, doctor.Surname, doctor.CRM, userID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("No rows were updated, doctor not found")
	}

	return nil
}

func (s *DoctorService) Delete(userID int, db *sql.DB) error {
	query := "DELETE FROM doctor WHERE user_id = ?"
	_, err := db.Exec(query, userID)

	if err != nil {
		return err
	}

	return nil
}
