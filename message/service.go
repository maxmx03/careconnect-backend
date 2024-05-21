package message

import "database/sql"

type MessageService struct{}

func (s *MessageService) GetAll(doctorID int, patientID int, db *sql.DB) ([]MessageModel, error) {
	var messages []MessageModel
	query := `
    SELECT 
        m.message_id,
        m.datetime,
        m.content,
        sender.user_type AS sender_type,
        COALESCE(sender_patient.name, sender_doctor.name) AS sender_name,
        COALESCE(sender_patient.surname, sender_doctor.surname) AS sender_surname,
        recipient.user_type AS recipient_type,
        COALESCE(recipient_patient.name, recipient_doctor.name) AS recipient_name,
        COALESCE(recipient_patient.surname, recipient_doctor.surname) AS recipient_surname
    FROM 
        message m
    INNER JOIN 
        user sender ON m.sender_id = sender.user_id
    LEFT JOIN 
        patient sender_patient ON sender.user_id = sender_patient.user_id
    LEFT JOIN 
        doctor sender_doctor ON sender.user_id = sender_doctor.user_id
    INNER JOIN 
        user recipient ON m.recipient_id = recipient.user_id
    LEFT JOIN 
        patient recipient_patient ON recipient.user_id = recipient_patient.user_id
    LEFT JOIN 
        doctor recipient_doctor ON recipient.user_id = recipient_doctor.user_id
    WHERE 
        (sender.user_id = ? AND sender.user_type = 'doctor') OR 
        (recipient.user_id = ? AND recipient.user_type = 'patient');
    `

	rows, err := db.Query(query, doctorID, patientID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message MessageModel
		err := rows.Scan(
			&message.MessageID,
			&message.DateTime,
			&message.Content,
			&message.SenderType,
			&message.SenderName,
			&message.SenderSurname,
			&message.RecipientType,
			&message.RecipientName,
			&message.RecipientSurname,
		)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return messages, nil
}

func (s *MessageService) Create(message *MessageModel, db *sql.DB) error {
	query := `
    INSERT INTO message (sender_id, recipient_id, datetime, content)
    VALUES (?, ?, ?, ?)
    `
	_, err := db.Exec(query, message.SenderID, message.RecipientID, message.DateTime, message.Content)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) Update(message *MessageModel, db *sql.DB) error {
	query := `
    UPDATE message
    SET sender_id = ?, recipient_id = ?, datetime = ?, content = ?
    WHERE message_id = ?
    `
	_, err := db.Exec(query, message.SenderID, message.RecipientID, message.DateTime, message.Content, message.MessageID)
	if err != nil {
		return err
	}
	return nil
}

func (s *MessageService) Delete(messageID int, db *sql.DB) error {
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
