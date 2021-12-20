package domain

import (
	"database/sql"
	"time"

	"github.com/Striker87/Banking/errs"
	"github.com/Striker87/Banking/logger"
	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {
	rows, err := d.client.Query("SELECT customer_id, `name`, city, zipcode, date_of_birth, `status` FROM customers")
	if err != nil {
		logger.Error("error while querying customer table due error: " + err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("error while scanning customer due error: " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer
	err := d.client.QueryRow("SELECT customer_id, `name`, city, zipcode, date_of_birth, `status` FROM customers WHERE customer_id = ?", id).Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
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
	client, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client}
}
