package medicalprescription

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var medicalPrescriptionController = &MedicalPrescriptionController{}

func Routes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
	var url = "/prescriptions"
	g := e.Group(url)
	g.Use(m...)

	g.GET("", func(c echo.Context) error {
		return medicalPrescriptionController.GetAll(c, db)
	})
	g.POST("", func(c echo.Context) error {
		return medicalPrescriptionController.Create(c, db)
	})
	g.PUT("/:id", func(c echo.Context) error {
		return medicalPrescriptionController.Update(c, db)
	})
	g.DELETE("/:id", func(c echo.Context) error {
		return medicalPrescriptionController.Delete(c, db)
	})
}
