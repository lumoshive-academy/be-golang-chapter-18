// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// type User struct {
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

// 	// Username dan password yang diterima dari input pengguna (misalnya dari form login)
// 	username := "admin"
// 	password := "password' OR '1'='1"

// 	// Query yang tidak aman
// 	// query := "SELECT id, username, password, email FROM customers WHERE username='" + username + "' AND password='" + password + "'"
// 	query := fmt.Sprintf("SELECT id, username, password, email FROM customers WHERE username='%s' AND password='%s'", username, password)
// 	fmt.Println("Executing query:", query)

// 	row := db.QueryRow(query)

// 	var user User
// 	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			fmt.Println("No user found with the provided credentials.")
// 		} else {
// 			log.Fatal(err)
// 		}
// 	} else {
// 		fmt.Printf("User found: %+v\n", user)
// 	}
// }

package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Customer struct {
	ID       int
	Username string
	Password string
	Email    string
}

func main() {
	// Koneksi ke database
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Input pengguna yang aman dari SQL Injection
	username := "admin"
	password := "adminpassword"

	// Query yang aman menggunakan parameterized query
	query := "SELECT id, username, password, email FROM customers WHERE username=$1 AND password=$2"
	fmt.Println("Executing query:", query)

	row := db.QueryRow(query, username, password)

	var customer Customer
	err = row.Scan(&customer.ID, &customer.Username, &customer.Password, &customer.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No customer found with the provided credentials.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Printf("Customer found: %+v\n", customer)
	}
}
