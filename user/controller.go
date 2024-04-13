package user

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserController struct{}

var userService UserRepository = &UserService{}

func (u *UserController) GetUsers(c echo.Context, db *sql.DB) error {
	users, err := userService.GetUsers(db)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserController) CreateUser(c echo.Context, db *sql.DB) error {

	user := &UserModel{}

	if err := c.Bind(u); err != nil {
		return err
	}

	err := userService.CreateUser(user, db)

  if err != nil {
    return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
  }

	return c.JSON(http.StatusCreated, u)
}
