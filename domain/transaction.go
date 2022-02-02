package domain

import "github.com/Striker87/Banking/dto"

const WITHDRAWAL = "withdrawal"

type Transaction struct {
	AccountId       string  `db:"account_id"`
	Amount          float64 `db:"amount"`
	TransactionDate string  `db:"transaction_date"`
	TransactionId   string  `db:"transaction_id"`
	TransactionType string  `db:"transaction_type"`
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == WITHDRAWAL
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionDate: t.TransactionDate,
		TransactionId:   t.TransactionId,
		TransactionType: t.TransactionType,
	}
}
