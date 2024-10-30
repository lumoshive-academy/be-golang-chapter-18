package repository

import (
	"database/sql"
	"golang-pertemuan-18/model"
)

type CustomerRepositoryDB struct {
	DB *sql.DB
}

func NewCustomerRepository(db *sql.DB) CustomerRepository {
	return &CustomerRepositoryDB{DB: db}
}

func (r *CustomerRepositoryDB) Create(customer *model.Customer) error {
	query := `INSERT INTO customer (username, password) VALUES ($1, $2) RETERNING id`
	err := r.DB.QueryRow(query, customer.Username, customer.Password).Scan(customer.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepositoryDB) GetAll(customer *[]model.Customer) (*model.Customer, error) {
	return nil, nil
}
