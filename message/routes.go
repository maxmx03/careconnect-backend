package message

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"github.com/maxmx03/careconnect-backend/token"
	"net/http"
)

var messageController = &MessageController{}

func DoctorRoutes(e *echo.Echo, db *sql.DB, m ...echo.MiddlewareFunc) {
	e.GET("/message", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return messageController.GetMessages(c, db)
	}, m...)

	e.POST("/message", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return messageController.CreateMessage(c, db)
	}, m...)

	e.PUT("/message/:id", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return messageController.UpdateMessage(c, db)
	}, m...)

	e.DELETE("/message/:id", func(c echo.Context) error {
		if err := token.ValidateToken(c); err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return messageController.DeleteMessage(c, db)
	}, m...)
}
