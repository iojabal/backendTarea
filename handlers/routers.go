package handlers

import (
	"backend/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RouterInit(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1:5500", "http://localhost:5173"},
		AllowMethods:     []string{echo.POST},
		AllowHeaders:     []string{echo.HeaderContentType},
		AllowCredentials: true,
	}))

	users := e.Group("/users")
	users.Use(middleware.CORSWithConfig(middleware.CORSConfig{AllowOrigins: []string{"*"}}))
	users.POST("/register", controllers.RegisterUser)
	users.POST("/login", controllers.LoginUser)
	users.GET("/:id", controllers.GetAllUsers)
	users.DELETE("/:id", controllers.DeleteUser)
	users.PUT("/update", controllers.UpdateUser)
}
