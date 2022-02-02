package service

import (
	"time"

	"github.com/Striker87/Banking/domain"
	"github.com/Striker87/Banking/dto"
	"github.com/Striker87/Banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	MakeTransaction(request dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountType: req.AccountType,
		Amount:      req.Amount,
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(dbTSLayout),
		Status:      "1",
	}

	newAccount, err := s.repo.Save(account)
	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAccountResponseDto()

	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTransactionTypewithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionDate: time.Now().Format(dbTSLayout),
		TransactionType: req.TransactionType,
	}
	transaction, appError := s.repo.SaveTransaction(t)
	if appError != nil {
		return nil, appError
	}
	response := transaction.ToDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repo}
}
