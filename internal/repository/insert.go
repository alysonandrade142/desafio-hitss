package repository

import (
	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/db"
)

func Insert(user model.User) (id int64, err error) {

	conn, err := db.GetConnection()
	if err != nil {
		return
	}

	defer conn.Close()

	sql := `INSERT INTO users (nome, sobrenome, contato, endereço, data_nasc, cpf) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err = conn.QueryRow(sql, user.Nome, user.Sobrenome, user.Contato, user.Endereço, user.DataNasc, user.CPF).Scan(&id)

	return
}
