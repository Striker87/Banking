package dto

import "github.com/Striker87/Banking/errs"

const (
	DEPOSIT    = "deposit"
	WITHDRAWAL = "withdrawal"
)

type TransactionRequest struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"amount"`
	CustomerId      string  `json:"-"`
	TransactionDate string  `json:"transaction_date"`
	TransactionType string  `json:"transaction_type"`
}

type TransactionResponse struct {
	AccountId       string  `json:"account_id"`
	Amount          float64 `json:"new_balance"`
	TransactionDate string  `json:"transaction_date"`
	TransactionId   string  `json:"transaction_id"`
	TransactionType string  `json:"transaction_type"`
}

func (r TransactionRequest) IsTransactionTypewithdrawal() bool {
	return r.TransactionType == WITHDRAWAL
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.TransactionType != WITHDRAWAL && r.TransactionType != DEPOSIT {
		return errs.NewValidationError("Transaction type can only be deposit or withdrawal")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount cannot be less than zero")
	}
	return nil
}
