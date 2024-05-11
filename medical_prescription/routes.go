package medicalprescription

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/maxmx03/careconnect-backend/token"
	"net/http"
)

var medicalPrescriptionController = &MedicalPrescriptionController{}

func DoctorRoutes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
	e.GET("/medicalprescription", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return medicalPrescriptionController.GetMedicalPrescriptionsById(c, db)
	}, m...)

	e.POST("/medicalprescription", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return medicalPrescriptionController.CreateMedicalPrescription(c, db)
	}, m...)

	e.PUT("/medicalprescription/:id", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return medicalPrescriptionController.UpdateMedicalPrescription(c, db)
	}, m...)

	e.DELETE("/medicalprescription/:id", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return medicalPrescriptionController.DeleteMedicalPrescription(c, db)
	}, m...)
}
