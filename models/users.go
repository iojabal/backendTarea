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
	query := "SELECT id, name, lastname, username, password FROM users WHERE username = ? LIMIT 1;"
	row := db.QueryRow(query, u.Username)
	return row, nil
}

func FetchUsersDB(id string) (*sql.Rows, error) {
	db, err := config.Connection()
	if err != nil {
		return nil, fmt.Errorf(err.Error())

	}
	defer db.Close()
	query := "SELECT username, name, lastname FROM users WHERE id = ?"
	rows, err := db.Query(query, id)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	return rows, nil
}

func FetchUserDB(id string) (*sql.Row, error) {
	db, err := config.Connection()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	defer db.Close()
	query := "SELECT username, name, lastname, password FROM users WHERE id = ?"
	row := db.QueryRow(query, id)
	return row, nil

}

func DeleteUserDB(id string) error {
	db, err := config.Connection()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer db.Close()
	query := "DELETE FROM users WHERE id = ?"
	_, err = db.Exec(query, id)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	return nil
}

func (u *Users) UpdateUserDB() error {
	db, err := config.Connection()
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	defer db.Close()
	query := "UPDATE users SET name = ?, lastname = ?, username = ?, password = ? WHERE id = ?;"
	_, err = db.Exec(query, u.Name, u.LastName, u.Username, u.Password, u.Id)
	if err != nil {
		return err
	}
	return nil
}
