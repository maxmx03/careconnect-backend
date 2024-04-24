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
	"github.com/maxmx03/careconnect-backend/token"
	"github.com/maxmx03/careconnect-backend/user"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://localhost:4200"},
	}))
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

	auth.AuthRoutes(e, db)
	user.UserRoutes(e, db, jwtMiddleware)
	doctor.DoctorRoutes(e, db)
	e.Logger.Fatal(e.Start(":3000"))
}
