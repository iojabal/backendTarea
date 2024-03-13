package services

import (
	"backend/models"

	"github.com/golang-jwt/jwt/v5"
)

func SignedLoginToken(u *models.Users) (string, *models.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     u.Name,
		"lastname": u.LastName,
		"username": u.Username,
	})

	jwtString, err := token.SignedString([]byte("Secret"))
	if err != nil {
		return "", &models.Error{Error: err, Type: "jwt"}
	}
	return jwtString, nil
}
