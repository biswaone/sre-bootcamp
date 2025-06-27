package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() {
	connStr := os.Getenv("DATABASE_URL")
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}
}

func GetDB() *sql.DB {
	return db
}
