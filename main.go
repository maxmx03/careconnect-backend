package main

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/maxmx03/careconnect-backend/auth"
	"github.com/maxmx03/careconnect-backend/doctor"
	"github.com/maxmx03/careconnect-backend/medical_prescription"
	"github.com/maxmx03/careconnect-backend/message"
	"github.com/maxmx03/careconnect-backend/patient"
	"github.com/maxmx03/careconnect-backend/token"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORS())

	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/careconnect")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	privateKey, err := os.ReadFile("ec256-private.pem")

	if err != nil {
		panic(err)
	}

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(token.JwtCustomClaims)
		},
		SigningKey: privateKey,
	}

	jwtMiddleware := echojwt.WithConfig(config)

	auth.Routes(e, db)
	doctor.Routes(e, db, jwtMiddleware)
	patient.Routes(e, db, jwtMiddleware)
	message.Routes(e, db, jwtMiddleware)
	medicalprescription.Routes(e, db, jwtMiddleware)
	e.Logger.Fatal(e.Start(":3000"))
}
