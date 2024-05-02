package patient

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/message"
)

type PatientController struct{}

var patientService PatientRepository = &PatientService{}

func (u *PatientController) GetPatients(c echo.Context, db *sql.DB) error {
	patients, err := patientService.GetPatients(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusFound, GetError("Failed to fetch patients"))
	}

	return c.JSON(http.StatusOK, patients)
}

func (u *PatientController) GetPatientById(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusFound, GetError("Failed to fetch doctors"))
	}

	var patient *PatientModel

	if patient, err = patientService.GetPatientById(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusFound, GetError("Failed to fetch patients"))
	}

	return c.JSON(http.StatusOK, patient)
}

func (u *PatientController) CreatePatient(c echo.Context, db *sql.DB) error {
	patient := &PatientModel{}

	if err := c.Bind(patient); err != nil {
		return err
	}

	err := patientService.CreatePatient(patient, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, GetError("Failed to create patient"))
	}

	return c.JSON(http.StatusCreated, patient)
}

func (u *PatientController) UpdatePatient(c echo.Context, db *sql.DB) error {
	patient := &PatientModel{}
	var err error

	if err := c.Bind(patient); err != nil {
		return err
	}

	err = patientService.UpdatePatient(patient, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, GetError("Failed to update patient"))
	}

	return c.JSON(http.StatusOK, GetOk("Patient updated successfully"))
}

func (u *PatientController) DeletePatient(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, GetError("Failed to fetch doctors"))
	}

	if err = patientService.DeletePatient(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, GetError("Failed to fetch patients"))
	}

	return c.JSON(http.StatusOK, GetOk("Patient deleted successfully"))
}
