package domain

import (
	"database/sql"
	"time"

	"github.com/Striker87/Banking/errs"
	"github.com/Striker87/Banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var err error
	var customers = make([]Customer, 0)

	if status == "" {
		err = d.client.Select(&customers, "SELECT customer_id, `name`, city, zipcode, date_of_birth, `status` FROM customers")
	} else {
		err = d.client.Select(&customers, "SELECT customer_id, `name`, city, zipcode, date_of_birth, `status` FROM customers WHERE status = ?", status)
	}

	if err != nil {
		logger.Error("error while querying customer table due error: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpectred database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	err := d.client.Get(&c, "SELECT customer_id, `name`, city, zipcode, date_of_birth, `status` FROM customers WHERE customer_id = ?", id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("error while scanning customer by id due error: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	client, err := sqlx.Connect("mysql", "root:@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
