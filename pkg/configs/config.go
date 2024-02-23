package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var conf *config

type config struct {
	API *APIConfig
	DB  *DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() {
	defaultValues := map[string]interface{}{
		"api.port":      "9000",
		"database.host": "localhost",
		"database.port": "5432",
		"database.user": "postgres",
		"database.pass": "",
		"database.name": "",
	}
	for key, value := range defaultValues {
		viper.SetDefault(key, value)
	}
}

func Load() error {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../pkg/configs/")

	err := viper.ReadInConfig()
	if err != nil {
		println(err.Error())
	}

	conf = &config{
		API: &APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: &DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetString("database.port"),
			User:     viper.GetString("database.user"),
			Pass:     viper.GetString("database.pass"),
			Database: viper.GetString("database.name"),
		},
	}

	fmt.Printf("Loaded config: %+v\n", conf.API)

	return nil
}

func GetDB() *DBConfig {
	return conf.DB
}

func GetServerPort() string {
	return conf.API.Port
}
