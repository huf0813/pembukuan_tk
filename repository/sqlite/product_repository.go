package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
	"time"
)

type ProductRepo struct {
	SqlConn sqlite.ConnSqlite
}

func (pr *ProductRepo) GetProducts() ([]entity.ProductStock, error) {
	conn := pr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	queryStock := "(select sum(product_increases.quantity) from product_increases where product_increases.product_id=products.id) stock"
	queryQtyInvoice := "(select sum(product_decreases.quantity) from product_decreases where product_decreases.product_id=products.id) qty_invoice"
	stringQuery := fmt.Sprintf("select id, name, price, %s, %s from products", queryStock, queryQtyInvoice)
	rows, err := conn.Query(stringQuery)
	if err != nil {
		return nil, err
	}

	var result []entity.ProductStock
	for rows.Next() {
		var dataRowProduct entity.ProductStock
		var qtyProduct sql.NullInt64
		var qtyInvoice sql.NullInt64
		if err := rows.Scan(&dataRowProduct.ID,
			&dataRowProduct.Name,
			&dataRowProduct.Price,
			&qtyProduct, &qtyInvoice); err != nil {
			return nil, err
		}
		if qtyProduct.Valid {
			dataRowProduct.Stock += qtyProduct.Int64
		}
		if qtyInvoice.Valid {
			dataRowProduct.Stock -= qtyInvoice.Int64
		}
		result = append(result, dataRowProduct)
	}
	return result, nil
}

func (pr *ProductRepo) AddProduct(newUser *entity.Product) (*entity.Product, error) {
	conn := pr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into products(name, price) values (?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(newUser.Name, newUser.Price)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.Product{
		ID:    int(lastInsertedID),
		Name:  newUser.Name,
		Price: newUser.Price,
	}, nil
}

func (pr *ProductRepo) EditProductByID(editedProduct *entity.Product) (*entity.Product, error) {
	conn := pr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("update products set name=?, price=? where id=?")
	if err != nil {
		return nil, err
	}
	if _, err := result.Exec(editedProduct.Name, editedProduct.Price, editedProduct.ID); err != nil {
		return nil, err
	}

	return &entity.Product{
		ID:    editedProduct.ID,
		Name:  editedProduct.Name,
		Price: editedProduct.Price,
	}, nil
}

func (pr *ProductRepo) AddProductStock(addQuantity *entity.ProductIncrease) (*entity.ProductIncrease, error) {
	conn := pr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into product_increases(product_id, quantity, user_id, created_at, updated_at) values (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(addQuantity.ProductID, addQuantity.Quantity, addQuantity.UserID, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.ProductIncrease{
		ID:        int(lastInsertedID),
		ProductID: addQuantity.ProductID,
		Quantity:  addQuantity.Quantity,
		UserID:    addQuantity.UserID,
	}, nil
}

func (pr *ProductRepo) DecProductStock(decQuantity *entity.ProductDec) (*entity.ProductDec, error) {
	conn := pr.SqlConn.SqliteConn()
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
