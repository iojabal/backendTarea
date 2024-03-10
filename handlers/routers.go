package handlers

import (
	"backend/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RouterInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.POST("/register", controllers.RegisterUser)
	e.POST("/login", controllers.LoginUser)
}
