package medicalconsultation

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/feedback"
	"net/http"
	"strconv"
)

type MedicalConsultationController struct{}

var medicalConsultationService MedicalConsultationRepository = &MedicalConsultationService{}

func (m *MedicalConsultationController) GetAll(c echo.Context, db *sql.DB) error {

	doctorID, err := strconv.Atoi(c.QueryParam("doctor_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid doctor id"))
	}

	patientID, err := strconv.Atoi(c.QueryParam("patient_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid patient id"))
	}

	medicalConsultations, err := medicalConsultationService.GetAll(doctorID, patientID, db)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("medical consultation not found"))
	}

	return c.JSON(http.StatusOK, medicalConsultations)
}

func (m *MedicalConsultationController) Create(c echo.Context, db *sql.DB) error {
	var medicalConsultation *MedicalConsultationModel
	if err := c.Bind(medicalConsultation); err != nil {
		return err
	}

	if err := medicalConsultationService.Create(medicalConsultation, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to create medical consultation"))
	}

	return c.JSON(http.StatusCreated, GetOk("Medical consultation created successfully"))
}

func (m *MedicalConsultationController) Update(c echo.Context, db *sql.DB) error {
	var medicalConsultation *MedicalConsultationModel
	if err := c.Bind(medicalConsultation); err != nil {
		return err
	}

	if err := medicalConsultationService.Update(medicalConsultation, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to update medical prescription"))
	}

	return c.JSON(http.StatusOK, GetOk("Medical consultation updated successfully"))
}

func (m *MedicalConsultationController) Delete(c echo.Context, db *sql.DB) error {
	consultationID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, GetError("Invalid consultation ID"))
	}

	if err := medicalConsultationService.Delete(consultationID, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to delete medical consultation"))
	}

	return c.JSON(http.StatusOK, GetOk("Medical prescription deleted successfully"))
}
