package services

import (
	"backend/models"
)

func RegisterClientService(client *models.Client) error {
	err := client.RegisterClientDB()
	if err != nil {
		return err
	}
	return nil
}

func FetchClientService(client *models.Client) ([]models.Client, *models.Error) {
	usersRows, err := client.FetchUSersDB()
	if err != nil {
		return nil, &models.Error{Error: err, Type: "db"}
	}
	var clients []models.Client
	for usersRows.Next() {
		var clientF models.Client
		err := usersRows.Scan(
			&clientF.Id,
			&clientF.Ci,
			&clientF.Name,
			&clientF.LastName,
			&clientF.Email,
			&clientF.Phone,
			&clientF.Lat,
			&clientF.Long)
		if err != nil {
			return nil, &models.Error{Error: err, Type: "db"}
		}
		clients = append(clients, clientF)

		if err := usersRows.Err(); err != nil {
			return nil, &models.Error{Error: err, Type: "db"}
		}
	}
	return clients, nil

}
