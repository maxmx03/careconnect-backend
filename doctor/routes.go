package doctor

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/maxmx03/careconnect-backend/token"
	"net/http"
)

var doctorController = &DoctorController{}

func DoctorRoutes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
	e.GET("/doctors", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return doctorController.GetDoctors(c, db)
	}, m...)
	e.GET("/user", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return doctorController.GetDoctorById(c, db)
	}, m...)
	e.POST("/user", func(c echo.Context) error {
		return doctorController.CreateDoctor(c, db)
	})
	e.PUT("/user", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return doctorController.UpdateDoctor(c, db)
	}, m...)
	e.DELETE("/user", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return doctorController.DeleteDoctor(c, db)
	}, m...)
}
