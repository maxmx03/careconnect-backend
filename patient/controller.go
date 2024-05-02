package patient

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type PatientController struct{}

var patientService PatientRepository = &PatientService{}

func (u *PatientController) GetPatients(c echo.Context, db *sql.DB) error {
	users, err := patientService.GetPatients(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch patients"})
	}

	return c.JSON(http.StatusOK, users)
}

func (u *PatientController) GetPatientById(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch doctors"})
	}

	var user *PatientModel

	if user, err = patientService.GetPatientById(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch patients"})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *PatientController) CreatePatient(c echo.Context, db *sql.DB) error {
	user := &PatientModel{}

	if err := c.Bind(user); err != nil {
		return err
	}

	err := patientService.CreatePatient(user, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create patient"})
	}

	return c.JSON(http.StatusCreated, user)
}

func (u *PatientController) UpdatePatient(c echo.Context, db *sql.DB) error {
	user := &PatientModel{}
	var err error

	if err := c.Bind(user); err != nil {
		return err
	}

	err = patientService.UpdatePatient(user, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update patient"})
	}

	return c.JSON(http.StatusOK, "Patient updated successfully")
}

func (u *PatientController) DeletePatient(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch doctors"})
	}

	if err = patientService.DeletePatient(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch patients"})
	}

	return c.JSON(http.StatusOK, "Patient deleted successfully")
}
