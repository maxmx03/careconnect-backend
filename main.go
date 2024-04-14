package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/maxmx03/careconnect-backend/user"
)

func main() {
	e := echo.New()
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/careconnect")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	user.Routes(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}
