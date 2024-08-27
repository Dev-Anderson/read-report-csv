package main

import (
	"log"
	"read-report-csv/internal/database"
	"read-report-csv/internal/utils"
)

func main() {
	db, err := database.GetDatabase()
	if err != nil {
		log.Fatalln("Falha na conexao com o banco de dados", err)
	}
	defer db.Close()

	utils.ReadLineCSV(db)

}
