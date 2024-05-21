package patient

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var patientController = &PatientController{}

func Routes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
  var url = "/patients"
	g := e.Group(url)
	g.Use(m...)

	g.GET("", func(c echo.Context) error {
		return patientController.GetAll(c, db)
	})
	g.GET("/:id", func(c echo.Context) error {
		return patientController.GetById(c, db)
	})
	e.POST(url, func(c echo.Context) error {
		return patientController.Create(c, db)
	})
  g.PUT("/:id", func(c echo.Context) error {
		return patientController.Update(c, db)
	})
	g.DELETE("/:id", func(c echo.Context) error {
		return patientController.Delete(c, db)
	})
}
