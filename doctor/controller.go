package doctor

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type DoctorController struct{}

var doctorService DoctorRepository = &DoctorService{}

func (u *DoctorController) GetDoctors(c echo.Context, db *sql.DB) error {
	doctors, err := doctorService.GetDoctors(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch doctors"})
	}

	return c.JSON(http.StatusOK, doctors)
}

func (u *DoctorController) GetDoctorById(c echo.Context, db *sql.DB) error {
	doctor := &DoctorModel{}
	var err error

	if err := c.Bind(doctor); err != nil {
		return err
	}

	if doctor, err = doctorService.GetDoctorById(doctor, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch doctors"})
	}

	return c.JSON(http.StatusOK, doctor)
}

func (u *DoctorController) CreateDoctor(c echo.Context, db *sql.DB) error {
	doctor := &DoctorModel{}

	if err := c.Bind(doctor); err != nil {
		return err
	}

	err := doctorService.CreateDoctor(doctor, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create doctor"})
	}

	return c.JSON(http.StatusCreated, doctor)
}

func (u *DoctorController) UpdateDoctor(c echo.Context, db *sql.DB) error {
	doctor := &DoctorModel{}
	var err error

	if err := c.Bind(doctor); err != nil {
		return err
	}

	err = doctorService.UpdateDoctor(doctor, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update doctor"})
	}

	return c.JSON(http.StatusOK, "doctor updated successfully")
}

func (u *DoctorController) DeleteDoctor(c echo.Context, db *sql.DB) error {
	doctor := &DoctorModel{}
	var err error

	if err := c.Bind(doctor); err != nil {
		return err
	}

	if err = doctorService.DeleteDoctor(doctor, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch doctors"})
	}

	return c.JSON(http.StatusOK, "Doctor deleted successfully")
}
