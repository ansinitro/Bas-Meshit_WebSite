package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDB() error {
	var err error

	// Использование DATABASE_URL из переменных окружения
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	DB, err = sql.Open("postgres", databaseUrl)
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	return DB.Close()
}
