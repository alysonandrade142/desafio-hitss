package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("../../pkg/configs/")

	err := viper.ReadInConfig()
	if err != nil {
		println(err.Error())
	}

	fmt.Printf("Loaded config: %+v\n", viper.GetString("api.port"))
}
