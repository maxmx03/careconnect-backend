package doctor

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/message"
)

type DoctorController struct{}

var doctorService DoctorRepository = &DoctorService{}

func (u *DoctorController) GetDoctors(c echo.Context, db *sql.DB) error {
	doctors, err := doctorService.GetDoctors(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, GetError("Failed to fetch doctors"))
	}

	return c.JSON(http.StatusOK, doctors)
}

func (u *DoctorController) GetDoctorById(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Invalid doctor id"))
	}

	var doctor *DoctorModel

	if doctor, err = doctorService.GetDoctorById(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, GetError("Failed to fetch doctor"))
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
		return c.JSON(http.StatusInternalServerError, GetError("Failed to create doctor"))
	}

	return c.JSON(http.StatusCreated, doctor)
}

func (u *DoctorController) UpdateDoctor(c echo.Context, db *sql.DB) error {
	doctor := &DoctorModel{}

	if err := c.Bind(doctor); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Invalid doctor id"))
	}

	err = doctorService.UpdateDoctor(doctor, id, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to update doctor"))
	}

	return c.JSON(http.StatusOK, GetOk("doctor updated successfully"))
}

func (u *DoctorController) DeleteDoctor(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusInternalServerError, GetError("Invalid doctor id"))
	}

	if err = doctorService.DeleteDoctor(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, GetError("Failed to delete doctor"))
	}

	return c.JSON(http.StatusOK, GetOk("Doctor deleted successfully"))
}
