package services

import (
	"backend/models"
	"net/http"
)

func HandlerErrors(err *models.Error) (int, map[string]string) {
	switch err.Type {
	case "db", "jwt", "bs64":
		return http.StatusInternalServerError, map[string]string{"Message": "Error Interno"}
	case "user":
		return http.StatusUnauthorized, map[string]string{"Message": "Usuario o contraseña incorrecta"}
	default:
		return http.StatusOK, map[string]string{}
	}
}
