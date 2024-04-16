package doctor

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var doctorController = &DoctorController{}

func DoctorRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/doctor", func(c echo.Context) error {
		return doctorController.GetDoctor(c, db)
	})
}
