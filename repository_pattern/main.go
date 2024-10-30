package main

import (
	"golang-pertemuan-18/database"
	"golang-pertemuan-18/service"
	"log"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// connStr := "user=postgres password=postgres dbname=postgres sslmode=disable"
	// db, err := sql.Open("postgres", connStr)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// defer db.Close()

	service.InputDataCustomer(db, "lumoshive", "123")
}
