package db

import (
	"database/sql"
	"fmt"

	"github.com/alysonandrade142/desafio-hitss/pkg/configs"
	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	conf := configs.GetDB()

	sc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", conf.Host, conf.Port, conf.User, conf.Pass, conf.Database)

	db, err := sql.Open("postgres", sc)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, err
}
