package service

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-pertemuan-18/model"
	"golang-pertemuan-18/repository"
)

func InputDataCustomer(db *sql.DB, username string, password string) error {
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}

	customerRepo := repository.NewCustomerRepository(db)
	customer := model.Customer{
		Username: username,
		Password: password,
	}

	customerRepo.Create(&customer)

	fmt.Println("berhasil input data customer dengan id ", customer.ID)

	return nil
}
