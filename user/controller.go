package user

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	. "github.com/maxmx03/careconnect-backend/feedback"
)

type UserController struct{}

var userService UserRepository = &UserService{}

func (a *UserController) Login(c echo.Context, db *sql.DB) error {
	auth := &UserModel{}
	var err error
	var token string

	if err = c.Bind(auth); err != nil {
		return err
	}

	if token, err = userService.Login(auth, db); err != nil {
		return c.JSON(http.StatusForbidden, GetError(err.Error()))
	}

	return c.JSON(http.StatusOK, map[string]string{"token": token})
}

func (a *UserController) Create(c echo.Context, db *sql.DB) error {
	user := &UserModel{}

	if err := c.Bind(user); err != nil {
		return err
	}

  err := userService.Create(user, db)

  if err!= nil {
    log.Error(err)
    return c.JSON(http.StatusInternalServerError, GetError("Failed to create user"))
  }

  return c.JSON(http.StatusCreated, GetOk("User created successfully"))
}

func (u *UserController) Update(c echo.Context, db *sql.DB) error {
	user := &UserModel{}

	if err := c.Bind(user); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusNotFound, GetError("Invalid doctor id"))
	}

	err = userService.Update(user, id, db)

	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to update user"))
	}

	return c.JSON(http.StatusOK, GetOk("User updated successfully"))
}

func (u *UserController) Delete(c echo.Context, db *sql.DB) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, GetError("Invalid user id"))
	}

	if err = userService.Delete(id, db); err != nil {
		log.Error(err)
		return c.JSON(http.StatusNotFound, GetError("Failed to delete user"))
	}

	return c.JSON(http.StatusOK, GetOk("User deleted successfully"))
}
