package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	. "github.com/maxmx03/careconnect-backend/user"
)

func main() {
	e := echo.New()
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/careconnect")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	UserController := &UserController{}

	e.GET("/users", func(c echo.Context) error {
		return UserController.GetUsers(c, db)
	})
  e.GET("/user", func(c echo.Context) error {
    return UserController.GetUserById(c, db)
  })
	e.POST("/user", func(c echo.Context) error {
		return UserController.CreateUser(c, db)
	})
  e.DELETE("/user", func(c echo.Context) error {
return UserController.DeleteUser(c, db)
  })

	e.Logger.Fatal(e.Start(":3000"))
}
