package medicalprescription

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/feedback"
)

type MedicalPrescriptionController struct{}

var medicalPrescriptionService MedicalPrescriptionRepository = &MedicalPrescriptionService{}

func (m *MedicalPrescriptionController) GetAll(c echo.Context, db *sql.DB) error {
	doctorID, err := strconv.Atoi(c.QueryParam("doctor_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid doctor id"))
	}

	patientID, err := strconv.Atoi(c.QueryParam("patient_id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid patient id"))
	}

	medicalPrescriptions, err := medicalPrescriptionService.GetAll(doctorID, patientID, db)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to fetch medical prescriptions"))
	}

	return c.JSON(http.StatusOK, medicalPrescriptions)
}

func (m *MedicalPrescriptionController) Create(c echo.Context, db *sql.DB) error {
	var medicalPrescription *MedicalPrescriptionModel
	if err := c.Bind(medicalPrescription); err != nil {
		return err
	}

	if err := medicalPrescriptionService.Create(medicalPrescription, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to create medical prescription"))
	}

	return c.JSON(http.StatusCreated, GetOk("Medical prescription created successfully"))
}

func (m *MedicalPrescriptionController) Update(c echo.Context, db *sql.DB) error {
	var medicalPrescription *MedicalPrescriptionModel
	if err := c.Bind(medicalPrescription); err != nil {
		return err
	}

	if err := medicalPrescriptionService.Update(medicalPrescription, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to update medical prescription"))
	}

	return c.JSON(http.StatusOK, GetOk("Medical prescription updated successfully"))
}

func (m *MedicalPrescriptionController) Delete(c echo.Context, db *sql.DB) error {
	prescriptionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, GetError("Invalid prescription ID"))
	}

	if err := medicalPrescriptionService.Delete(prescriptionID, db); err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Failed to delete medical prescription"))
	}

	return c.JSON(http.StatusOK, GetOk("Medical prescription deleted successfully"))
}
