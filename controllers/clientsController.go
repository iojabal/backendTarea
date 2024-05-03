package controllers

import (
	"backend/models"
	"backend/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterClients(c echo.Context) error {
	client := new(models.Client)
	if err := c.Bind(client); err != nil {
		return err
	}

	err := services.RegisterClientService(client)

	if err != nil {
		return c.JSON(services.HandlerErrors(&models.Error{Error: err, Type: "db"}))
	}
	return c.JSON(http.StatusCreated, client)
}

func FetchAllClients(c echo.Context) error {
	// Create a new instance of Client
	client := &models.Client{}

	// Fetch clients from the database using FetchClientService
	clients, err := services.FetchClientService(client)
	if err != nil {
		// Return error response if there's an error
		return c.JSON(services.HandlerErrors(err))
	}

	// Return clients as JSON response
	return c.JSON(http.StatusOK, clients)
}
