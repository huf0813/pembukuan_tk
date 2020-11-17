package sqlite

import "github.com/huf0813/pembukuan_tk/model"

type CustomerRepo struct {
	CustomerModel model.Customer
}

type CustomerRepoInterface interface {
	GetCustomers() (*[]model.Customer, error)
	AddCustomer(name, phone, email, address string) (*model.Customer, error)
}

var customers = []model.Customer{
	{1, "jodi", "1234567890", "jodi@gmail.com", "malang"},
}

func (cr *CustomerRepo) GetCustomers() (*[]model.Customer, error) {
	return &customers, nil
}

func (cr *CustomerRepo) AddCustomer(name, phone, email, address string) (*model.Customer, error) {
	newCustomer := cr.CustomerModel
	newCustomer.Name = name
	newCustomer.Phone = phone
	newCustomer.Email = email
	newCustomer.Address = address
	customers = append(customers, newCustomer)
	return &newCustomer, nil
}
