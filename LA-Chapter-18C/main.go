// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"time"

// 	_ "github.com/lib/pq"
// )

// type User struct {
// 	ID               int32
// 	FirstName        string
// 	LastName         string
// 	Email            string
// 	BirthDate        time.Time
// 	RegistrationDate time.Time
// }

// func main() {
// 	// Koneksi ke database
// 	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Query untuk mengambil data dari tabel users
// 	rows, err := db.Query("SELECT id, first_name, last_name, email, birth_date, registration_date FROM users")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	// Membaca hasil query dan memetakkannya ke struct User
// 	var users []User
// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.RegistrationDate)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		users = append(users, user)
// 	}
// 	// Memeriksa apakah ada error selama pembacaan
// 	if err := rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	// Menampilkan data pengguna
// 	for _, u := range users {
// 		fmt.Printf("ID: %d, Name: %s %s, Email: %s, Birth Date: %s, Registration Date: %s\n", u.ID, u.FirstName, u.LastName, u.Email, u.BirthDate.Format("2006-01-02"), u.RegistrationDate.Format("2006-01-02 15:04:05"))
// 	}
// }

package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	ID               int64
	Username         sql.NullString
	Age              sql.NullInt64
	Email            sql.NullString
	BirthDate        sql.NullTime
	RegistrationDate time.Time
}

func main() {
	// Koneksi ke database
	db, err := sql.Open("postgres", "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Query untuk mengambil semua data dari tabel users
	rows, err := db.Query("SELECT id, username, age, email, birth_date, registration_date FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Membaca hasil query dan memetakkannya ke struct User
	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Age, &user.Email, &user.BirthDate, &user.RegistrationDate)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// Memeriksa error selama pembacaan
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Menampilkan data pengguna
	for _, u := range users {
		fmt.Printf("ID: %d\n", u.ID)
		if u.Username.Valid {
			fmt.Printf("Username: %s\n", u.Username.String)
		} else {
			fmt.Println("Username: <NULL>")
		}
		if u.Age.Valid {
			fmt.Printf("Age: %d\n", u.Age.Int64)
		} else {
			fmt.Println("Age: <NULL>")
		}

		if u.Email.Valid {
			fmt.Printf("Email: %s\n", u.Email.String)
		} else {
			fmt.Println("Email: <NULL>")
		}
		if u.BirthDate.Valid {
			fmt.Printf("Birth Date: %s\n", u.BirthDate.Time.Format("2006-01-02"))
		} else {
			fmt.Println("Birth Date: <NULL>")
		}
		fmt.Printf("Registration Date: %s\n", u.RegistrationDate.Format("2006-01-02 15:04:05"))
		fmt.Println("-------------------")
	}
}
