package database

import (
	"database/sql"
)

// var DB *sql.DB

func InitDB() (*sql.DB, error) {
	connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// DB = db
	// defer db.Close()
	return db, err
}
