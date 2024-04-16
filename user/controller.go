package user

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type UserController struct{}

var userService UserRepository = &UserService{}

func (u *UserController) GetUsers(c echo.Context, db *sql.DB) error {
	users, err := userService.GetUsers(db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserController) GetUserById(c echo.Context, db *sql.DB) error {
  user := &UserModel{}
	var err error

	if err := c.Bind(user); err != nil {
		return err
	}

	if user, err = userService.GetUserById(user, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserController) CreateUser(c echo.Context, db *sql.DB) error {
	user := &UserModel{}

	if err := c.Bind(user); err != nil {
		return err
	}

	err := userService.CreateUser(user, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(http.StatusCreated, user)
}

func (u *UserController) UpdateUser(c echo.Context, db *sql.DB) error {
  user := &UserModel{}
	var err error

	if err := c.Bind(user); err != nil {
		return err
	}

	err = userService.UpdateUser(user, user, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update user"})
	}

	return c.JSON(http.StatusOK, "User updated successfully")
}

func (u *UserController) DeleteUser(c echo.Context, db *sql.DB) error {
  user := &UserModel{}
	var err error

	if err := c.Bind(user); err != nil {
		return err
	}

	if err = userService.DeleteUser(user, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch users"})
	}

	return c.JSON(http.StatusOK, "User deleted successfully")
}
