package doctor

import (
	"database/sql"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/feedback"
)

type DoctorController struct{}

var doctorService DoctorRepository = &DoctorService{}

func (u *DoctorController) GetAll(c echo.Context, db *sql.DB) error {
	doctors, err := doctorService.GetAll(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusFound, GetError("Failed to fetch doctors"))
	}

	return c.JSON(http.StatusOK, doctors)
}

func (u *DoctorController) GetById(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid doctor id"))
	}

	var doctor *DoctorModel

	if doctor, err = doctorService.GetById(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to fetch doctor"))
	}

	return c.JSON(http.StatusOK, doctor)
}

func (u *DoctorController) Create(c echo.Context, db *sql.DB) error {
	doctor := &DoctorModel{}

	if err := c.Bind(doctor); err != nil {
		return err
	}

	err := doctorService.Create(doctor, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, GetError("Failed to create doctor"))
	}

	return c.JSON(http.StatusCreated, GetOk("Doctor created successfully"))
}

func (u *DoctorController) Update(c echo.Context, db *sql.DB) error {
	doctor := &DoctorModel{}

	if err := c.Bind(doctor); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid doctor id"))
	}

	err = doctorService.Update(doctor, id, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to update doctor"))
	}

	return c.JSON(http.StatusOK, GetOk("doctor updated successfully"))
}

func (u *DoctorController) Delete(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, GetError("Invalid doctor id"))
	}

	if err = doctorService.Delete(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to delete doctor"))
	}

	return c.JSON(http.StatusOK, GetOk("Doctor deleted successfully"))
}
