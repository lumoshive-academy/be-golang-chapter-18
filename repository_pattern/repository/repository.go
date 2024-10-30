package repository

import "golang-pertemuan-18/model"

type CustomerRepository interface {
	Create(customer *model.Customer) error
	GetAll(customer *[]model.Customer) (*model.Customer, error)
}
