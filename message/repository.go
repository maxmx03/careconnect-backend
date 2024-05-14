package message

import "database/sql"

type MessageRepository interface {
	GetMessages(doctorID int, patientID int, db *sql.DB) ([]MessageModel, error)
	CreateMessage(message *MessageModel, db *sql.DB) error
	UpdateMessage(message *MessageModel, db *sql.DB) error
	DeleteMessage(messageID int, db *sql.DB) error
}
