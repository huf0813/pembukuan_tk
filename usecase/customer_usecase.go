package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type CustomerUseCase struct {
	CustomerRepo sqlite.CustomerRepo
}

type CustomerUseCaseInterface interface {
	AddNewCustomer(name, phone, email, address string) (*entity.Customer, error)
	GetCustomers() ([]entity.Customer, error)
}

func (cuc *CustomerUseCase) GetCustomers() ([]entity.Customer, error) {
	result, err := cuc.CustomerRepo.GetCustomers()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cuc *CustomerUseCase) AddCustomerValidation(name, phone, email, address string) error {
	if name == "" || phone == "" || email == "" || address == "" {
		return errors.New("fields cannot be empty")
	}
	return nil
}

func (cuc *CustomerUseCase) AddCustomer(name, phone, email, address string) (*entity.Customer, error) {
	if err := cuc.AddCustomerValidation(name, phone, email, address); err != nil {
		return nil, err
	}
	result, err := cuc.CustomerRepo.AddCustomer(name, phone, email, address)
	if err != nil {
		return nil, err
	}
	return result, nil
}
