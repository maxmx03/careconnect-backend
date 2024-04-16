package doctor

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/user"
)

type DoctorController struct{}

var doctorService DoctorRepository = &DoctorService{}

func (d *DoctorController) GetDoctor(c echo.Context, db *sql.DB) error {
 user := &UserModel{}
 doctor := &DoctorModel{}
 var err error

 if err = c.Bind(user); err != nil {
   return err
 }

	if doctor, err = doctorService.GetDoctor(user, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch doctor"})
	}

 return c.JSON(http.StatusOK, doctor)
}
