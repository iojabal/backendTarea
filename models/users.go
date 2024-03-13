package models

import (
	"backend/config"
	"database/sql"
	"fmt"
)

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *Users) RegisterUserDB() error {
	db, err := config.Connection()

	if err != nil {
		return err
	}
	defer db.Close()
	query := "INSERT INTO users(name, lastname, username, password) VALUES (?, ?, ?, ?);"
	_, err = db.Exec(query, u.Name, u.LastName, u.Username, u.Password)

	if err != nil {
		return err
	}
	return nil
}

func (u *Users) FetchUserDB() (*sql.Row, error) {
	db, err := config.Connection()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer db.Close()
	query := "SELECT name, lastname, username, password FROM users WHERE username = ? LIMIT 1;"
	row := db.QueryRow(query, u.Username)
	return row, nil
}
