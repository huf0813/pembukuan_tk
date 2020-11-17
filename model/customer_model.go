package model

type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func (c *Customer) NewCustomer(name, phone, email, address string) *Customer {
	return &Customer{
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address,
	}
}
