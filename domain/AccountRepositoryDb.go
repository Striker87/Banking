package domain

import (
	"strconv"

	"github.com/Striker87/Banking/errs"
	"github.com/Striker87/Banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	result, err := d.client.Exec("INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)", a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}

func NewAccountRepositoryDb(db *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{db}
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	result, _ := tx.Exec("INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) VALUES(?, ?, ?, ?)", t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	if t.IsWithdrawal() {
		_, err = tx.Exec("UPDATE accounts SET amount = amount - ? WHERE account_id = ?", t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec("UPDATE accounts SET amount = amount + ? WHERE account_id = ?", t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while savig transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commiting transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// getting the latestaccount informationfrom theaccounts table
	account, appErr := d.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating transaction struct with the latest balance
	t.Amount = account.Amount

	return &t, nil
}

func (d AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	var account Account
	err := d.client.Get(&account, "SELECT account_id, customer_id, opening_date, account_type, amount FROM `accounts` WHERE account_id = ?", accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &account, nil
}
