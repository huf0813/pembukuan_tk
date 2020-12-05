package usecase

import (
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type CustomerUseCase struct {
	CustomerRepo sqlite.CustomerRepo
}

type CustomerUseCaseInterface interface {
	AddNewCustomer(name, phone, email, address string) (*entity.Customer, error)
	GetCustomers() (*[]entity.Customer, error)
}

func (cuc *CustomerUseCase) GetCustomers() ([]entity.Customer, error) {
	result, err := cuc.CustomerRepo.GetCustomers()
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (cuc *CustomerUseCase) AddNewCustomer(name, phone, email, address string) (*entity.Customer, error) {
	result, err := cuc.CustomerRepo.AddCustomer(name, phone, email, address)
	if err != nil {
		return nil, err
	}
	return result, nil
}
