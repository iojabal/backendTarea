package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
	u := new(models.Users)
	if err := c.Bind(u); err != nil {
		return err
	}

	services.RegisterUserService(u)

	return c.JSON(http.StatusCreated, u)
}

func LoginUser(c echo.Context) error {
	u := new(models.Users)
	if err := c.Bind(u); err != nil {
		return err
	}

	err := services.LoginUserService(u)
	switch err != nil {

	case err.Type == "db":
		return c.JSON(http.StatusInternalServerError, map[string]string{"Message": "Error Interno"})
	case err.Type == "user":
		return c.JSON(http.StatusUnauthorized, map[string]string{"Message": "Usuario o contraseña incorrecta"})
	}
	return c.JSON(http.StatusOK, map[string]string{"Message": "Bienvenido Usuario"})
}
