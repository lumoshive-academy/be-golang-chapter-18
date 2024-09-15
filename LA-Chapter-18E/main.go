// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// type Customer struct {
// 	ID       int
// 	Username string
// 	Password string
// 	Email    string
// }

// func main() {
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Menyiapkan statement SQL
// 	stmt, err := db.Prepare("SELECT id, username, password, email FROM customers WHERE username=$1 AND password=$2")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// 	// Mengeksekusi statement dengan parameter
// 	username := "admin"
// 	password := "adminpassword"

// 	row := stmt.QueryRow(username, password)

// 	var customer Customer
// 	err = row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No customer found with the provided credentials.")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		fmt.Printf("Customer found: %+v\n", customer)
// 	}

// 	// Mengeksekusi statement dengan parameter yang berbeda
// 	username = "customer1"
// 	password = "password1"

// 	row = stmt.QueryRow(username, password)

// 	err = row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No customer found with the provided credentials.")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		fmt.Printf("Customer found: %+v\n", customer)
// 	}
// }

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	// Koneksi ke database
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// untuk mendapatkan ID dari baris yang dimasukkan
	var lastInsertId int
	err = db.QueryRow("INSERT INTO customers (username, password, email) VALUES ($1, $2, $3) RETURNING id",
		"dikcy", "123456", "dicky@example.com").Scan(&lastInsertId)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Last Inserted ID: %d\n", lastInsertId)
}
