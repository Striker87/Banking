package domain

import (
	"github.com/Striker87/Banking/dto"
	"github.com/Striker87/Banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	Status      string  `db:"status"`
}

type AccountRepository interface {
	FindBy(accountId string) (*Account, *errs.AppError)
	Save(account Account) (*Account, *errs.AppError)
	SaveTransaction(transaction Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountId: a.AccountId}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount < amount
}
