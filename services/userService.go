package services

import (
	"backend/models"
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUserService(users *models.Users) error {

	users.Password, _ = EncryptPassword(users.Password)
	users.Password = EncodeBase64(users.Password)
	users.RegisterUserDB()

	return nil
}

func LoginUserService(user *models.Users) *models.Error {
	var hashedPassword string
	row, err := user.FetchUserDB()

	if err != nil {
		return &models.Error{Error: err, Type: "db"}
	}
	row.Scan(&user.Id, &user.Name, &user.LastName, &user.Username, &hashedPassword)
	hashedPassword, err = DecodeBase64(hashedPassword)
	if err != nil {
		return &models.Error{Error: err, Type: "bs64"}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		return &models.Error{Error: err, Type: "user"}
	}
	return nil

}

func FetchUser(id string, user *models.Users) *models.Error {
	row, err := models.FetchUserDB(id)
	if err := row.Scan(&user.Name, &user.LastName, &user.Username, &user.Password); err != nil {
		return &models.Error{Error: err, Type: "db"}
	}
	if err != nil {
		return &models.Error{Error: err, Type: "db"}
	}
	return nil
}

func FetchAllUsers(id string) (*sql.Rows, *models.Error) {
	rows, err := models.FetchUsersDB(id)
	if err != nil {
		return nil, &models.Error{Error: err, Type: "db"}
	}
	return rows, nil
}

func DeleteUser(id string) *models.Error {
	err := models.DeleteUserDB(id)
	if err != nil {
		return &models.Error{Error: err, Type: "db"}
	}
	return nil
}

func UpdateUser(user *models.Users) *models.Error {
	err := user.UpdateUserDB()

	if err != nil {
		return &models.Error{Error: err, Type: "db"}
	}
	return nil
}
