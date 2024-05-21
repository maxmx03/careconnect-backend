package message

import "database/sql"

type MessageRepository interface {
	GetAll(doctorID int, patientID int, db *sql.DB) ([]MessageModel, error)
	Create(message *MessageModel, db *sql.DB) error
	Update(message *MessageModel, db *sql.DB) error
	Delete(messageID int, db *sql.DB) error
}
