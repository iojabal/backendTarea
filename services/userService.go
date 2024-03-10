package services

import (
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserService(users *models.Users) error {

	users.Password, _ = EncryptPassword(users.Password)
	users.RegisterUserDB()

	return nil
}

func LoginUserService(user *models.Users) *models.Error {
	var hashedPassword string
	row, err := user.FetchUserDB()

	if err != nil {
		return &models.Error{Error: err, Type: "db"}
	}
	row.Scan(&user.Username, &hashedPassword)
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		return &models.Error{Error: err, Type: "user"}
	}
	return nil

}
