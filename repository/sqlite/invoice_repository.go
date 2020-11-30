package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/model"
)

type InvoiceRepo struct {
	SqlConn sqlite.ConnSqlite
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
