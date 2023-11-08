package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func OpenDB() error {
	var err error
	const (
		host     = "localhost"
		port     = 7777
		user     = "postgres"
		password = "15691804"
		dbname   = "BasMeshit"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	DB, err = sql.Open("postgres", connStr)

	if err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	return DB.Close()
}
