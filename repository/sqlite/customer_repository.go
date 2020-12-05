package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
)

type CustomerRepo struct {
	CustomerModel entity.Customer
	SqlConn       sqlite.ConnSqlite
}

type CustomerRepoInterface interface {
	GetCustomers() ([]entity.Customer, error)
	AddCustomer(name, phone, email, address string) (*entity.Customer, error)
}

func (cr *CustomerRepo) GetCustomers() ([]entity.Customer, error) {
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

	var result []entity.Customer
	for rows.Next() {
		var dataRow entity.Customer
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

func (cr *CustomerRepo) AddCustomer(name, phone, email, address string) (*entity.Customer, error) {
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

	return &entity.Customer{
		ID:      int(lastInsertedID),
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address,
	}, nil
}

func (cr *CustomerRepo) EditCustomer(name, phone, email, address string, customerID int) (*entity.Customer, error) {
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
		conn.Prepare("update customers set name=?, phone=?, email=?, address=? where id=?")
	if err != nil {
		return nil, err
	}
	if _, err := result.Exec(name, phone, email, address, customerID); err != nil {
		return nil, err
	}

	return &entity.Customer{
		ID:      customerID,
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address,
	}, nil
}
