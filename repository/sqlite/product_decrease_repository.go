package sqlite

import (
	"errors"
	"fmt"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
)

type ProductDecreaseRepo struct {
	SqlConn sqlite.ConnSqlite
}

func (pdr *ProductDecreaseRepo) GetProductDecreases() ([]entity.ProductInsideInvoice, error) {
	conn := pdr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	// get products inside invoice
	stringQueryGetProductsInsideInvoice := "SELECT products.name, products.price, product_decreases.quantity, (products.price * product_decreases.quantity) product_total_price from products join product_decreases on products.id=product_decreases.product_id where product_decreases.invoice_id=?"
	rows, err := conn.Query(stringQueryGetProductsInsideInvoice)
	if err != nil {
		return nil, err
	}

	var result []entity.ProductInsideInvoice
	for rows.Next() {
		var dataRow entity.ProductInsideInvoice
		if err := rows.Scan(&dataRow.ProductName,
			&dataRow.ProductPrice,
			&dataRow.ProductQty,
			&dataRow.ProductTotalPrice); err != nil {
			return nil, err
		}
		result = append(result, dataRow)
	}
	return result, nil
}

func (pdr *ProductDecreaseRepo) GetProductDecreaseByID(invoiceID int) ([]entity.ProductInsideInvoice, error) {
	conn := pdr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	// get products inside invoice
	stringQuery := fmt.Sprintf("SELECT products.name, products.price, product_decreases.quantity, (products.price * product_decreases.quantity) product_total_price from products join product_decreases on products.id=product_decreases.product_id where product_decreases.invoice_id=%d", invoiceID)
	rows, err := conn.Query(stringQuery)
	if err != nil {
		return nil, err
	}

	var result []entity.ProductInsideInvoice
	for rows.Next() {
		var dataRow entity.ProductInsideInvoice
		if err := rows.Scan(&dataRow.ProductName,
			&dataRow.ProductPrice,
			&dataRow.ProductQty,
			&dataRow.ProductTotalPrice); err != nil {
			return nil, err
		}
		result = append(result, dataRow)
	}
	return result, nil
}

func (pdr *ProductDecreaseRepo) AddDecProductStock(decQuantity *entity.ProductDec) (*entity.ProductDec, error) {
	conn := pdr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into product_decreases(product_id, quantity, invoice_id) values (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(decQuantity.ProductID, decQuantity.Quantity, decQuantity.InvoiceID)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.ProductDec{
		ID:        int(lastInsertedID),
		ProductID: decQuantity.ProductID,
		Quantity:  decQuantity.Quantity,
		InvoiceID: decQuantity.InvoiceID,
	}, nil
}
