package repository

import "github.com/alysonandrade142/desafio-hitss/pkg/db"

func Delete(id int64) (int64, error) {

	conn, err := db.GetConnection()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
