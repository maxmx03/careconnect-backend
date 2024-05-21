package medicalconsultation

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var medicalConsultationController = &MedicalConsultationController{}

func Routes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
	var url = "/consultations"
	g := e.Group(url)
	g.Use(m...)

	g.GET("", func(c echo.Context) error {
		return medicalConsultationController.GetAll(c, db)
	})
	g.POST("", func(c echo.Context) error {
		return medicalConsultationController.Create(c, db)
	})
	g.PUT("/:id", func(c echo.Context) error {
		return medicalConsultationController.Update(c, db)
	})
	g.DELETE("/:id", func(c echo.Context) error {
		return medicalConsultationController.Delete(c, db)
	})
}
