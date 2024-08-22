package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=1234 dbname=teste sslmode=disable")
	if err != nil {
		panic(err)
	}

	// defer db.Close()

	file, err := os.Open("assets/relatorio.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ','       //definicao do delimitador
	reader.LazyQuotes = true // lida com aspas dentro dos campo

	lines, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//intera sobre as linhas do CSV
	for _, line := range lines[1:] { //ignora o cabecalho

		query := `
			INSERT INTO relatorio (card_id, historical_board, historical_workflow, historical_section, historical_column, column_type, historical_column_parent, cycle_time_days, card_title, size_tamanho, last_moved, type_tipo, owners, last_modified, total_cycle_time)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15)
		`

		_, err := db.Exec(query, line[0], line[1], line[2], line[3], line[4], line[5], line[6], line[7], line[8], line[9], line[10], line[11], line[12], line[13], line[14])
		if err != nil {
			log.Printf("Erro ao inserir linha: %v, erro: %v\n", line, err)
		}
	}

	fmt.Println("Todos os dados inseridos com sucesso")

}
