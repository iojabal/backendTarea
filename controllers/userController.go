package controllers

import (
	"backend/models"
	"backend/services"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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

func GetAllUsers(c echo.Context) error {
	id := c.Param("id")

	rows, err := services.FetchAllUsers(id)
	if err != nil {
		return c.JSON(services.HandlerErrors(err))
	}
	var users []models.Users // Suponiendo que User es tu estructura de datos para un usuario

	for rows.Next() {
		var user models.Users                                        // Suponiendo que User es tu estructura de datos para un usuario
		err := rows.Scan(&user.Username, &user.Name, &user.LastName) // Suponiendo que tienes campos ID, Name y Email en tu estructura User
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "error escaneando fila"})
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": "error al iterar sobre filas"})
	}

	return c.JSON(http.StatusOK, users)
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")
	err := services.DeleteUser(id)
	if err != nil {
		return c.JSON(services.HandlerErrors(err))
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Usuario Eliminado"})
}

// func UpdateUser(c echo.Context) error {
// 	u := new(models.Users)
// 	if err := c.Bind(u); err != nil {
// 		return err
// 	}
// 	token := c.Request().Header["Cookie"]
// 	tokenString := strings.Split(token[0], "Bearer ")
// 	tokenParts, err := services.VerifyToken(tokenString[1])
// 	fmt.Println(tokenParts["id"])
// 	if err != nil {
// 		return c.JSON(services.HandlerErrors(err))
// 	}
// 	id := int(tokenParts["id"].(float64)) // Convert float64 to int
// 	u.Id = id
// 	u.UpdateUserDB()
// 	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Usuario Actualizado"})
// }

func UpdateUser(c echo.Context) error {
	token := c.Request().Header["Cookie"]
	tokenString := strings.Split(token[0], "Bearer ")
	tokenParts, err := services.VerifyToken(tokenString[1])
	id := int(tokenParts["id"].(float64))

	if err != nil {
		return c.JSON(services.HandlerErrors(err))
	}

	u := new(models.Users)
	user := &models.Users{}
	if err := services.FetchUser(strconv.Itoa(id), user); err != nil {
		return fmt.Errorf("error en algo")
	}
	fmt.Println(user)
	if err := c.Bind(u); err != nil {
		return fmt.Errorf("error en algo")
	}
	fmt.Print(u)
	return nil
}
