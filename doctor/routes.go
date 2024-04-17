package doctor

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/maxmx03/careconnect-backend/token"
)

var doctorController = &DoctorController{}

func DoctorRoutes(e *echo.Echo, db *sql.DB) {
	e.GET("/doctor", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, err)
		}

		return doctorController.GetDoctor(c, db)
	})
}
