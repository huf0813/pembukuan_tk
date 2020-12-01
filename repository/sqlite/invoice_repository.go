package sqlite

import (
	"errors"
	"fmt"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/model"
)

type InvoiceRepo struct {
	SqlConn        sqlite.ConnSqlite
	ProductDecRepo ProductDecreaseRepo
}

func (ir *InvoiceRepo) AddInvoice(newInvoice *model.Invoice) (*model.Invoice, error) {
	conn := ir.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into invoices(customer_id, user_id) values (?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(newInvoice.CustomerID, newInvoice.UserID)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.Invoice{
		ID:         int(lastInsertedID),
		CustomerID: newInvoice.CustomerID,
		UserID:     newInvoice.UserID,
	}, nil
}

func (ir *InvoiceRepo) GetInvoices() ([]model.InvoiceWithDetail, error) {
	conn := ir.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	// get invoices
	subQueryTotalPerInvoice := "(SELECT sum(product_decreases.quantity * products.price) from product_decreases join products on products.id=product_decreases.product_id where product_decreases.invoice_id=invoices.id) total_price"
	stringQuery := fmt.Sprintf("SELECT invoices.id, customers.name, customers.phone, customers.email, customers.address, %s from invoices join customers on customers.id=invoices.customer_id", subQueryTotalPerInvoice)
	rows, err := conn.Query(stringQuery)
	if err != nil {
		return nil, err
	}

	var result []model.InvoiceWithDetail
	for rows.Next() {
		var dataRow model.InvoiceWithDetail
		if err := rows.Scan(&dataRow.ID,
			&dataRow.CustomerName,
			&dataRow.CustomerPhone,
			&dataRow.CustomerEmail,
			&dataRow.CustomerAddress,
			&dataRow.TotalInvoicePrice); err != nil {
			return nil, err
		}
		if dataRow.Products, err = ir.ProductDecRepo.GetProductDecreaseByID(dataRow.ID); err != nil {
			return nil, err
		}
		result = append(result, dataRow)
	}
	return result, nil
}
