package message

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/feedback"
)

type MessageController struct{}

var messageService MessageRepository = &MessageService{}

func (m *MessageController) GetMessages(c echo.Context, db *sql.DB) error {
	doctorID, err := strconv.Atoi(c.QueryParam("doctor_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid doctor id"))
	}

	patientID, err := strconv.Atoi(c.QueryParam("patient_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid patient id"))
	}

	messages, err := messageService.GetMessages(doctorID, patientID, db)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to fetch messages"))
	}

	return c.JSON(http.StatusOK, messages)
}

func (m *MessageController) CreateMessage(c echo.Context, db *sql.DB) error {
	var message *MessageModel
	if err := c.Bind(message); err != nil {
		return err
	}

	if err := messageService.CreateMessage(message, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to create message"))
	}

	return c.JSON(http.StatusCreated, GetOk("Message created successfully"))
}

func (m *MessageController) UpdateMessage(c echo.Context, db *sql.DB) error {
	var medicalPrescription *MessageModel
	if err := c.Bind(medicalPrescription); err != nil {
		return err
	}

	if err := messageService.UpdateMessage(medicalPrescription, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to update medical prescription"))
	}

	return c.JSON(http.StatusOK, GetOk("Medical prescription updated successfully"))
}

func (m *MessageController) DeleteMessage(c echo.Context, db *sql.DB) error {
	prescriptionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, GetError("Invalid prescription ID"))
	}

	if err := messageService.DeleteMessage(prescriptionID, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to delete medical prescription"))
	}

	return c.JSON(http.StatusOK, GetOk("Medical prescription deleted successfully"))
}
