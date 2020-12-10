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
	conn := pr.SqlConn.SqliteConnInit()
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
	stringQuery := fmt.Sprintf("select id, name, price, %s, %s from products where deleted_at is null", queryStock, queryQtyInvoice)
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
	conn := pr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into products(name, price, created_at, updated_at) values (?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(newUser.Name, newUser.Price, time.Now().Unix(), time.Now().Unix())
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
	conn := pr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("update products set name=?, price=?, updated_at=? where id=?")
	if err != nil {
		return nil, err
	}
	edited, err := result.Exec(editedProduct.Name, editedProduct.Price, time.Now().Unix(), editedProduct.ID)
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

	return &entity.Product{
		ID:    editedProduct.ID,
		Name:  editedProduct.Name,
		Price: editedProduct.Price,
	}, nil
}

func (pr *ProductRepo) DeleteProductByID(productID int) (string, error) {
	conn := pr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return "", errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("update products set deleted_at=? where id=?")
	if err != nil {
		return "", err
	}
	edited, err := result.Exec(time.Now().Unix(), productID)
	if err != nil {
		return "", err
	}
	effected, err := edited.RowsAffected()
	if err != nil {
		return "", err
	}
	if effected == 0 {
		return "", errors.New("no data edited")
	}

	return "product deleted successfully", nil
}
