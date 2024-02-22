package repository

import (
	"github.com/alysonandrade142/desafio-hitss/internal/model"
	"github.com/alysonandrade142/desafio-hitss/pkg/db"
)

func Update(id int64, user model.User) (int64, error) {

	conn, err := db.GetConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE users SET nome = $1, sobrenome = $2, contato = $3, endereço = $4, data_nasc = $5, cpf = $6 WHERE id = $7`,
		user.Nome, user.Sobrenome, user.Contato, user.Endereço, user.DataNasc, user.CPF, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
