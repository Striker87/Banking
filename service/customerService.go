package service

import (
	"github.com/Striker87/Banking/domain"
	"github.com/Striker87/Banking/dto"
	"github.com/Striker87/Banking/errs"
)

type CustomerService interface {
	GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}

	var response = make([]dto.CustomerResponse, len(c))
	for i, v := range c {
		response[i] = v.ToDto()
	}

	return response, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
