package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error)  {
	connStr := "user=postgres dbname=db_tour_destination sslmode=disable password=admin host=localhost"
	db, err := sql.Open("postgres", connStr)

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatal(err)
	}

	return db, err
}