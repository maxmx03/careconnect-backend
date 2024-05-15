package patient

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/maxmx03/careconnect-backend/token"
	"net/http"
)

var patientController = &PatientController{}

func Routes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
	e.GET("/patients", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return patientController.GetPatients(c, db)
	}, m...)
	e.GET("/patient/:id", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return patientController.GetPatientById(c, db)
	}, m...)
	e.POST("/patient", func(c echo.Context) error {
		return patientController.CreatePatient(c, db)
	})
	e.PUT("/patient", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return patientController.UpdatePatient(c, db)
	}, m...)
	e.DELETE("/patient/:id", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return patientController.DeletePatient(c, db)
	}, m...)
}
