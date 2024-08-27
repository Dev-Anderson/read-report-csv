package utils

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	config "read-report-csv/infra/env"
)

func openFile(nameFile string) (*os.File, error) {
	file, err := os.Open(nameFile)
	if err != nil {
		log.Fatalln("Falha ao abrir o arquivo", err)
	}

	return file, err
}

func ReadFile(delimitador rune) ([][]string, error) {
	e, err := config.GetEnv()
	if err != nil {
		log.Fatalln("Falha ao abrir o arquivo de configuracao", err)
	}

	o, err := openFile(e.NameFile)
	if err != nil {
		log.Fatalln("Falha ao abrir o arquivo", err)
	}
	defer o.Close()

	r := csv.NewReader(o)
	r.Comma = delimitador // definir um delimitador
	r.LazyQuotes = true   // lida com as aspas duplas dentro dos campos

	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	return lines, err
}

func ReadLineCSV(db *sql.DB) {
	lines, err := ReadFile(',')
	if err != nil {
		log.Fatalln("Falha ao ler o arquivo", err)
	}

	for _, line := range lines[1:] {
		query := `
		INSERT INTO gestaofrotas (card_id, historical_board, historical_workflow, historical_section, historical_column, size_tamanho, type_tipo, setor, co_owners, last_modified)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

		_, err := db.Exec(query, line[0], line[1], line[2], line[3], line[4], line[5], line[6], line[7], line[8], line[9])
		if err != nil {
			log.Printf("Erro ao inserir a linha: %v, error: %v\n", line, err)
		}

		fmt.Println("Dados inseridos com sucesso")
	}
}
