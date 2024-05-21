package doctor

import (
	"database/sql"
	"github.com/labstack/echo/v4"
)

var doctorController = &DoctorController{}

func Routes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
  var url = "/doctors"
	g := e.Group(url)
	g.Use(m...)

	g.GET("", func(c echo.Context) error {
		return doctorController.GetAll(c, db)
	})
	g.GET("/:id", func(c echo.Context) error {
		return doctorController.GetById(c, db)
	})
	e.POST(url, func(c echo.Context) error {
		return doctorController.Create(c, db)
	})
	g.PUT("/:id", func(c echo.Context) error {
		return doctorController.Update(c, db)
	})
	g.DELETE("/:id", func(c echo.Context) error {
		return doctorController.Delete(c, db)
	})
}
