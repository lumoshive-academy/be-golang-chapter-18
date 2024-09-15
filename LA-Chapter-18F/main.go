// database transaction
// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/lib/pq"
// )

// // OrderItem mewakili item dalam pesanan

// type OrderItem struct {
// 	ProductID int
// 	Quantity  int
// 	Price     float64
// }

// func main() {

// 	// Koneksi ke database PostgreSQL
// 	connStr := "user=postgres dbname=postgres sslmode=disable password=postgres host=localhost"
// 	db, err := sql.Open("postgres", connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()

// 	// Membuat pesanan
// 	userID := 1
// 	orderItems := []OrderItem{
// 		{ProductID: 1, Quantity: 2, Price: 50.00},
// 		{ProductID: 2, Quantity: 1, Price: 30.00},
// 	}

// 	orderID, err := createOrder(db, userID, orderItems)
// 	if err != nil {

// 		log.Fatal(err)

// 	}
// 	fmt.Printf("Order created successfully with ID: %d\n", orderID)

// }

// func createOrder(db *sql.DB, userID int, items []OrderItem) (int, error) {
// 	// Memulai transaksi
// 	tx, err := db.Begin()
// 	if err != nil {
// 		return 0, err
// 	}

// 	// Membuat entri baru dalam tabel orders
// 	var orderID int
// 	err = tx.QueryRow("INSERT INTO orders (user_id) VALUES ($1) RETURNING id", userID).Scan(&orderID)
// 	if err != nil {
// 		tx.Rollback()
// 		return 0, err
// 	}

// 	// Memproses setiap item dalam pesanan
// 	for _, item := range items {
// 		// Memperbarui stok produk
// 		_, err := tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2 AND stock >= $1", item.Quantity, item.ProductID)
// 		if err != nil {
// 			tx.Rollback()
// 			return 0, err
// 		}

// 		// Menambahkan item ke tabel order_items
// 		_, err = tx.Exec("INSERT INTO order_items (order_id, product_id, quantity, price) VALUES ($1, $2, $3, $4)", orderID, item.ProductID, item.Quantity, item.Price)
// 		if err != nil {
// 			tx.Rollback()
// 			return 0, err
// 		}

// 	}

// 	// Menyelesaikan transaksi
// 	err = tx.Commit()
// 	if err != nil {
// 		return 0, err
// 	}

// 	return orderID, nil
// }

// repository pattern
package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// User merupakan entitas yang akan disimpan di database
type User struct {
	ID               int
	FirstName        string
	LastName         string
	Email            string
	BirthDate        time.Time
	RegistrationDate time.Time
}

// UserRepository merupakan antarmuka untuk mengakses data User
type UserRepository interface {
	Create(user *User) error
	Update(user *User) error
	Delete(id int) error
	GetByID(id int) (*User, error)
	GetAll() ([]*User, error)
}

// UserRepositoryDB adalah implementasi dari UserRepository menggunakan database SQL
type UserRepositoryDB struct {
	DB *sql.DB
}

// Create akan membuat user baru di database
func (r *UserRepositoryDB) Create(user *User) error {
	query := `INSERT INTO users (first_name, last_name, email, birth_date, registration_date)
              VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := r.DB.QueryRow(query, user.FirstName, user.LastName, user.Email, user.BirthDate, user.RegistrationDate).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Update akan memperbarui data user di database berdasarkan ID
func (r *UserRepositoryDB) Update(user *User) error {
	query := `UPDATE users SET first_name=$1, last_name=$2, email=$3, birth_date=$4, registration_date=$5 WHERE id=$6`
	_, err := r.DB.Exec(query, user.FirstName, user.LastName, user.Email, user.BirthDate, user.RegistrationDate, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// Delete akan menghapus data user dari database berdasarkan ID
func (r *UserRepositoryDB) Delete(id int) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

// GetByID akan mengembalikan data user dari database berdasarkan ID
func (r *UserRepositoryDB) GetByID(id int) (*User, error) {
	query := "SELECT id, first_name, last_name, email, birth_date, registration_date FROM users WHERE id=$1"
	row := r.DB.QueryRow(query, id)
	user := &User{}
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.RegistrationDate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetAll akan mengembalikan semua data user dari database
func (r *UserRepositoryDB) GetAll() ([]*User, error) {
	query := "SELECT id, first_name, last_name, email, birth_date, registration_date FROM users"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []*User{}
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.BirthDate, &user.RegistrationDate)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Inisialisasi repository
func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepositoryDB{DB: db}
}

func main() {
	// Koneksi ke database
	connStr := "user=youruser password=yourpassword dbname=yourdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Menginisialisasi repository
	userRepo := NewUserRepository(db)

	// Contoh penggunaan repository
	// Membuat user baru
	user := &User{
		FirstName:        "John",
		LastName:         "Doe",
		Email:            "john.doe@example.com",
		BirthDate:        time.Date(1990, time.January, 1, 0, 0, 0, 0, time.UTC),
		RegistrationDate: time.Now(),
	}
	err = userRepo.Create(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User created successfully with ID: %d\n", user.ID)

	// Mendapatkan semua user
	users, err := userRepo.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Name: %s %s, Email: %s\n", u.ID, u.FirstName, u.LastName, u.Email)
	}

	// Mendapatkan user berdasarkan ID
	userByID, err := userRepo.GetByID(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User found by ID: %d - Name: %s %s, Email: %s\n", userByID.ID, userByID.FirstName, userByID.LastName, userByID.Email)

	// Menghapus user
	err = userRepo.Delete(user.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User with ID %d has been deleted\n", user.ID)
}
