package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
	"time"
)

type CustomerRepo struct {
	SqlConn sqlite.ConnSqlite
}

type CustomerRepoInterface interface {
	GetCustomers() ([]entity.Customer, error)
	AddCustomer(name, phone, email, address string) (*entity.Customer, error)
	EditCustomer(name, phone, email, address string, customerID int) (*entity.Customer, error)
	DeleteCustomer(customerID int) (string, error)
}

func (cr *CustomerRepo) GetCustomers() ([]entity.Customer, error) {
	conn := cr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	rows, err := conn.Query("select id, name, phone, email, address from customers where deleted_at is null")
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
	conn := cr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into customers(name, phone, email, address, created_at, updated_at) values (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(name, phone, email, address, time.Now().Unix(), time.Now().Unix())
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
	conn := cr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("update customers set name=?, phone=?, email=?, address=?, updated_at=? where id=? and deleted_at is null")
	if err != nil {
		return nil, err
	}
	edited, err := result.Exec(name, phone, email, address, time.Now().Unix(), customerID)
	if err != nil {
		return nil, err
	}
	effected, err := edited.RowsAffected()
	if err != nil {
		return nil, err
	}
	if effected == 0 {
		return nil, errors.New("no data edited")
	}

	return &entity.Customer{
		ID:      customerID,
		Name:    name,
		Phone:   phone,
		Email:   email,
		Address: address,
	}, nil
}

func (cr *CustomerRepo) DeleteCustomer(customerID int) (string, error) {
	conn := cr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return "", errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("update customers set email=?, phone=?, deleted_at=? where id=?")
	if err != nil {
		return "", err
	}
	edited, err := result.Exec(customerID, customerID, time.Now().Unix(), customerID)
	if err != nil {
		return "", err
	}
	effected, err := edited.RowsAffected()
	if err != nil {
		return "", err
	}
	if effected == 0 {
		return "", errors.New("no data deleted")
	}

	return "deleted successfully", nil
}
