package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config") // config/
	viper.SetConfigName("config") // config.yaml

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	serverHost := viper.GetString("server.host")
	serverPort := viper.GetString("server.port")
	fmt.Println(serverHost)
	fmt.Println(serverPort)
}
