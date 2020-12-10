package sqlite

import (
	"errors"
	"fmt"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
	"time"
)

type InvoiceRepo struct {
	SqlConn        sqlite.ConnSqlite
	ProductDecRepo ProductDecreaseRepo
}

func (ir *InvoiceRepo) AddInvoice(newInvoice *entity.Invoice) (*entity.Invoice, error) {
	conn := ir.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into invoices(customer_id, user_id, created_at, updated_at) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(newInvoice.CustomerID, newInvoice.UserID, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.Invoice{
		ID:         int(lastInsertedID),
		CustomerID: newInvoice.CustomerID,
		UserID:     newInvoice.UserID,
	}, nil
}

func (ir *InvoiceRepo) GetInvoices() ([]entity.InvoiceWithDetail, error) {
	conn := ir.SqlConn.SqliteConnInit()
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
	stringQuery := fmt.Sprintf("SELECT invoices.id, customers.name, customers.phone, customers.email, customers.address, invoices.created_at, invoices.updated_at, %s from invoices join customers on customers.id=invoices.customer_id where invoices.deleted_at is null", subQueryTotalPerInvoice)
	rows, err := conn.Query(stringQuery)
	if err != nil {
		return nil, err
	}

	var result []entity.InvoiceWithDetail
	for rows.Next() {
		var dataRow entity.InvoiceWithDetail
		if err := rows.Scan(&dataRow.ID,
			&dataRow.CustomerName,
			&dataRow.CustomerPhone,
			&dataRow.CustomerEmail,
			&dataRow.CustomerAddress,
			&dataRow.CreatedAt,
			&dataRow.UpdatedAt,
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

func (ir *InvoiceRepo) DeleteInvoiceByID(invoiceID int) (string, error) {
	conn := ir.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return "", errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("delete from invoices where id=?")
	if err != nil {
		return "", err
	}
	edited, err := result.Exec(invoiceID)
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

	result2, err :=
		conn.Prepare("delete from product_decreases where invoice_id=?")
	if err != nil {
		return "", err
	}
	edited2, err := result2.Exec(invoiceID)
	if err != nil {
		return "", err
	}
	effected2, err := edited2.RowsAffected()
	if err != nil {
		return "", err
	}
	if effected2 == 0 {
		return "", errors.New("no data deleted")
	}

	return "invoice deleted successfully", nil
}
