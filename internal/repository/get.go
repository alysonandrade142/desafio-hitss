package repository

import (
	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/db"
)

func Get(id int64) (user model.User, err error) {

	conn, err := db.GetConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM users WHERE id = $1`, id)
	err = row.Scan(&user.ID, &user.Nome, &user.Sobrenome, &user.Contato, &user.Endereço, &user.DataNasc, &user.CPF)

	return
}
