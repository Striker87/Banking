package domain

import "github.com/Striker87/Banking/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	Zipcode     string
	DateofBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(id string) (*Customer, *errs.AppError)
}
