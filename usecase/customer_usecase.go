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
	GetCustomers() ([]entity.Customer, error)
	AddCustomerValidation(name, phone, email, address string) error
	AddCustomer(name, phone, email, address string) (*entity.Customer, error)
	EditCustomerValidation(name, phone, email, address string, customerID int) error
	EditCustomer(name, phone, email, address string, customerID int) (*entity.Customer, error)
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

func (cuc *CustomerUseCase) EditCustomerValidation(name, phone, email, address string, customerID int) error {
	if name == "" || phone == "" || email == "" || address == "" || customerID == 0 {
		return errors.New("fields cannot be empty")
	}
	return nil
}

func (cuc *CustomerUseCase) EditCustomer(name, phone, email, address string, customerID int) (*entity.Customer, error) {
	if err := cuc.EditCustomerValidation(name, phone, email, address, customerID); err != nil {
		return nil, err
	}
	result, err := cuc.CustomerRepo.EditCustomer(name, phone, email, address, customerID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
