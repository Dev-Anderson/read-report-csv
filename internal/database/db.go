package database

import (
	"database/sql"
	"fmt"
	"log"
	config "read-report-csv/infra/env"

	_ "github.com/lib/pq"
)

func GetDatabase() (*sql.DB, error) {
	e, err := config.GetEnv()
	if err != nil {
		log.Fatalln("Falha ao carregar o arquivo de configuracao", err)
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", e.DBHost, e.DBPort, e.DBUser, e.DBPassword, e.DBName)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return db, nil

}
