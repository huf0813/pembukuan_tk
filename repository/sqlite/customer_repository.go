package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/model"
)

type CustomerRepo struct {
	CustomerModel model.Customer
	SqlConn       sqlite.ConnSqlite
}

type CustomerRepoInterface interface {
	GetCustomers() ([]model.Customer, error)
	AddCustomer(name, phone, email, address string) (*model.Customer, error)
}

func (cr *CustomerRepo) GetCustomers() ([]model.Customer, error) {
	conn := cr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	rows, err := conn.Query("select id, name, phone, email, address from customers")
	if err != nil {
		return nil, err
	}

	var result []model.Customer
	for rows.Next() {
		var dataRow model.Customer
		if err := rows.Scan(&dataRow.ID,
			&dataRow.Name,
			&dataRow.Phone,
			&dataRow.Email,
			&dataRow.Address); err != nil {
			return nil, err
		}
		result = append(result, dataRow)
	}

	return result, nil
}

func (cr *CustomerRepo) AddCustomer(name, phone, email, address string) (*model.Customer, error) {
	conn := cr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into customers(name, phone, email, address) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(name, phone, email, address)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.Customer{
		ID:      int(lastInsertedID),
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address,
	}, nil
}
