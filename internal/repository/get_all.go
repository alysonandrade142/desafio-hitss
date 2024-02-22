package repository

import (
	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/db"
)

func GetAll() (users []model.User, err error) {

	conn, err := db.GetConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM users`)
	if err != nil {
		return
	}

	// Using a loop to iterate through the rows (cursor) of a result set. For each row, it creates a new User instance,
	// scans the values from the row into the User instance, and then continues to the next row if there is an error.

	for rows.Next() {

		var user model.User

		err = rows.Scan(&user.ID, &user.Nome, &user.Sobrenome, &user.Contato, &user.Endereço, &user.DataNasc, &user.CPF)

		if err != nil {
			continue
		}
		users = append(users, user)
	}

	return
}
