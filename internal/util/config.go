package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
	Secret      string
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{
		Port:        os.Getenv("PORT"),
		DB_Name:     os.Getenv("MYSQL_DB"),
		DB_Username: os.Getenv("MYSQL_USER"),
		DB_Password: os.Getenv("MYSQL_PASSWORD"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		Secret:      os.Getenv("SECRET"),
	}

	return cfg
}
