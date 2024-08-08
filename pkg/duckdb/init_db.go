package duckdb

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	"github.com/AirtonLira/protobuf_realtime_notification/pkg/utils"

	_ "github.com/marcboeker/go-duckdb"
)

func InitDB() (*sql.DB, error) {
	projectRoot, err := utils.GetProjectRoot()
	if err != nil {
		log.Fatalf("failed to get project root: %v", err)
	}

	dbFile := filepath.Join(projectRoot, "notifications.db")

	// Verifica se o arquivo do banco de dados existe
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Printf("Database file %s does not exist, creating...\n", dbFile)
	} else {
		// Tenta abrir o banco de dados existente
		db, err := sql.Open("duckdb", dbFile)
		if err == nil {
			defer db.Close()
			// Verifica se a tabela notification existe para validar o banco de dados
			_, err = db.Exec("SELECT 1 FROM notification LIMIT 1;")
			if err == nil {
				log.Printf("Database file %s is valid.\n", dbFile)
				return db, nil
			} else {
				log.Printf("Database file %s is invalid, recreating...\n", dbFile)
				os.Remove(dbFile)
			}
		} else {
			log.Printf("Failed to open database file %s, recreating...\n", dbFile)
			os.Remove(dbFile)
		}
	}

	// Abre uma conexão com DuckDB (isso cria o arquivo se ele não existir)
	db, err := sql.Open("duckdb", dbFile)
	if err != nil {
		return nil, err
	}

	// Cria a tabela notification se ela não existir
	createTableQuery := `
    CREATE TABLE IF NOT EXISTS notification (
        id STRING,
        title STRING,
        message STRING,
        user_id STRING,
        timestamp BIGINT
    );
    `
	_, err = db.Exec(createTableQuery)
	if err != nil {
		db.Close()
		return nil, err
	}

	log.Println("Database and table are ready.")
	return db, nil
}
