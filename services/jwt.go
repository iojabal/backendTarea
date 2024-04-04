package services

import (
	"backend/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func SignedLoginToken(u *models.Users) (string, *models.Error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       u.Id,
		"name":     u.Name,
		"lastname": u.LastName,
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	jwtString, err := token.SignedString([]byte("Secret"))
	if err != nil {
		return "", &models.Error{Error: err, Type: "jwt"}
	}
	return jwtString, nil
}

func VerifyToken(token string) (jwt.MapClaims, *models.Error) {
	var p jwt.Parser
	t, _, err := p.ParseUnverified(token, jwt.MapClaims{})
	if err != nil {
		return nil, &models.Error{Error: err, Type: "jwt"}
	}
	tokenParts, ok := t.Claims.(jwt.MapClaims)

	if !ok {
		return nil, &models.Error{Error: fmt.Errorf("Error Extracting claims from token"), Type: "jwt"}
	}

	// fmt.Printf("test: %v\n", tokenParts["username"])
	return tokenParts, nil
}
