package models

import (
	"backend/config"
	"database/sql"
)

type Client struct {
	Id       int     `json:"id"`
	Ci       string  `json:"ci"`
	Name     string  `json:"name"`
	LastName string  `json:"lastname"`
	Email    string  `json:"email"`
	Phone    int32   `json:"phone"`
	Lat      float64 `json:"Lat"`
	Long     float64 `json:"Long"`
}

func (c *Client) RegisterClientDB() error {
	db, err := config.Connection()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO client(Ci, Name, LastName, Email, Phone, Lat, `Long`) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, c.Ci, c.Name, c.LastName, c.Email, c.Phone, c.Lat, c.Long)

	if err != nil {
		return err
	}
	return nil
}

func (c *Client) FetchUSersDB() (*sql.Rows, error) {
	db, err := config.Connection()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	query := "SELECT * FROM client"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
