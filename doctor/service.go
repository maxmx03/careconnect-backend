package doctor

import (
	"database/sql"
	. "github.com/maxmx03/careconnect-backend/user"
)

type DoctorService struct{}

func (s *DoctorService) GetDoctor(user *UserModel, db *sql.DB) (*DoctorModel, error) {
  doctor := &DoctorModel{}
	query := `
    SELECT
        doctor.*,
        user.name AS user_name,
        user.email AS user_email,
        user.type AS user_type
    FROM
        doctor
    INNER JOIN
        user
    ON
        doctor.user_id = user.user_id
    WHERE
        doctor.user_id = ?
  `
	err := db.QueryRow(query, user.User_id).Scan(&doctor.user_id, &doctor.crm, &doctor.doctor_id)

	if err != nil {
		return &DoctorModel{}, err
	}

	return doctor, nil
}
