package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Koneksi ke database
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// Insert data baru ke tabel users
	result, err := db.Exec("INSERT INTO users (first_name, last_name, email, birth_date, registration_date) VALUES ('Budi', 'Santoso', 'budi@example.com', '1990-05-15', NOW())")
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Rows inserted: %d\n", rowsAffected)

	// call function get all user
	GetAllUser(db)
}

// get all user
func GetAllUser(db *sql.DB) {
	// Mengambil data dari tabel users dengan Query
	rows, err := db.Query("SELECT id, first_name, last_name, email, birth_date, registration_date FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Membaca hasil query
	for rows.Next() {
		var id int
		var firstName, lastName, email string
		var birthDate, registrationDate sql.NullTime

		err := rows.Scan(&id, &firstName, &lastName, &email, &birthDate, &registrationDate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d\n", id)
		fmt.Printf("Nama: %s %s\n", firstName, lastName)
		fmt.Printf("Email: %s\n", email)
		if birthDate.Valid {
			fmt.Printf("Tanggal Lahir: %s\n", birthDate.Time.Format("2006-01-02"))
		} else {
			fmt.Println("Tanggal Lahir: <NULL>")
		}
		if registrationDate.Valid {
			fmt.Printf("Tanggal Pendaftaran: %s\n", registrationDate.Time.Format("2006-01-02 15:04:05"))
		} else {
			fmt.Println("Tanggal Pendaftaran: <NULL>")
		}
		fmt.Println("-------------------")
	}

	// Periksa apakah ada error selama proses pembacaan
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

}
