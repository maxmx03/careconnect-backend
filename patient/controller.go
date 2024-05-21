package patient

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/feedback"
)

type PatientController struct{}

var patientService PatientRepository = &PatientService{}

func (u *PatientController) GetAll(c echo.Context, db *sql.DB) error {
	patients, err := patientService.GetAll(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to fetch patients"))
	}

	return c.JSON(http.StatusOK, patients)
}

func (u *PatientController) GetById(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Failed to fetch patients"))
	}

	var patient *PatientModel

	if patient, err = patientService.GetById(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to fetch patients"))
	}

	return c.JSON(http.StatusOK, patient)
}

func (u *PatientController) Create(c echo.Context, db *sql.DB) error {
	patient := &PatientModel{}

	if err := c.Bind(patient); err != nil {
		return err
	}

	err := patientService.Create(patient, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, GetError("Failed to create patient"))
	}

	return c.JSON(http.StatusCreated, GetOk("Patient created successfully"))
}

func (u *PatientController) Update(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Failed to fetch patients"))
	}

	patient := &PatientModel{}

	if err := c.Bind(patient); err != nil {
		return err
	}

	err = patientService.Update(patient, id, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, GetError("Failed to update patient"))
	}

	return c.JSON(http.StatusOK, GetOk("Patient updated successfully"))
}

func (u *PatientController) Delete(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, GetError("Failed to fetch doctors"))
	}

	if err = patientService.Delete(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, GetError("Failed to fetch patients"))
	}

	return c.JSON(http.StatusOK, GetOk("Patient deleted successfully"))
}
