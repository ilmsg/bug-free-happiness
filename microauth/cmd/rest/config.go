package main

import "os"

type config struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
}

func loadConfig() *config {
	return &config{
		DbHost:     os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}
