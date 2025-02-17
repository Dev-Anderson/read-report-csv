package config

import (
	"os"

	"github.com/joho/godotenv"
)

type ConfigEnv struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	NameFile   string
}

func GetEnv() (ConfigEnv, error) {
	godotenv.Load()
	return ConfigEnv{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		NameFile:   os.Getenv("NAME_FILE"),
	}, nil
}
