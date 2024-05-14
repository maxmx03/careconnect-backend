package message

import "database/sql"

type MessageService struct{}

func (s *MessageService) GetMessages(doctorID int, patientID int, db *sql.DB) ([]MessageModel, error) {
	var messages []MessageModel
	query := `
SELECT
    message.message_id,
    message.datetime,
    message.content,
    d.doctor_id,
    d.name AS doctor_name,
    d.surname AS doctor_surname,
    p.patient_id,
    p.name AS patient_name,
    p.surname AS patient_surname
FROM
    message
INNER JOIN 
    doctor d ON message.doctor_id = d.doctor_id
INNER JOIN 
    patient p ON message.patient_id = p.patient_id
WHERE 
    d.doctor_id = ? AND p.patient_id = ?
  `
	rows, err := db.Query(query, doctorID, patientID)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var message MessageModel

		if err := rows.Scan(&message.DoctorID, &message.DateTime, &message.Content, &message.DoctorID,
			&message.DoctorName, &message.DoctorSurname, &message.PatientID,
			&message.PatientName, &message.PatientSurname); err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

func (s *MessageService) CreateMessage(message *MessageModel, db *sql.DB) error {
	query := `
    INSERT INTO message (doctor_id, patient_id, datetime, content)
    VALUES (?, ?, ?, ?)
    `
	_, err := db.Exec(query, message.DoctorID, message.PatientID, message.DateTime, message.Content)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) UpdateMessage(message *MessageModel, db *sql.DB) error {
	query := `
    UPDATE message
    SET datetime = ?, content = ?
    WHERE message_id = ?
    `
	_, err := db.Exec(query, message.DateTime, message.Content, message.MessageID)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) DeleteMessage(messageID int, db *sql.DB) error {
	query := `
    DELETE FROM message
    WHERE message_id = ?
    `
	_, err := db.Exec(query, messageID)
	if err != nil {
		return err
	}
	return nil
}
