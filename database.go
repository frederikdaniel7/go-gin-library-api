package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var db *sql.DB

func InitDB() error {
	var err error
	db, err = ConnectDB()
	if err!= nil {
        log.Fatalf("error connecting to database: %s", err.Error())
		return err
    }
	return err
}
func ConnectDB() (*sql.DB, error) {
	// "host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}