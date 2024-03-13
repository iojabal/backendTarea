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
	var token string
	if err := services.LoginUserService(u); err != nil {
		return c.JSON(services.HandlerErrors(err))
	}
	token, err := services.SignedLoginToken(u)
	if err != nil {
		return c.JSON(services.HandlerErrors(err))
	}

	return c.JSON(http.StatusOK, map[string]string{"Message": "Bienvenido Usuario", "token": token})
}
