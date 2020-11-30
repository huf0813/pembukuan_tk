package sqlite

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/model"
)

type ProductRepo struct {
	SqlConn sqlite.ConnSqlite
}

func (pr *ProductRepo) GetProducts() ([]model.ProductStockAndType, error) {
	conn := pr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}
	subQueryStock := "((select sum(product_increases.quantity) from product_increases where product_increases.product_id=products.id) - (select sum(product_decreases.quantity) from product_decreases where product_decreases.product_id=products.id)) stock"
	stringQuery := fmt.Sprintf("select id, name, price, %s from products", subQueryStock)
	rows, err := conn.Query(stringQuery)
	if err != nil {
		return nil, err
	}

	var result []model.ProductStockAndType
	for rows.Next() {
		var dataRowProduct model.ProductStockAndType
		var dataRowProductStock sql.NullInt64
		if err := rows.Scan(&dataRowProduct.ID,
			&dataRowProduct.Name,
			&dataRowProduct.Price,
			&dataRowProductStock); err != nil {
			return nil, err
		}
		if dataRowProductStock.Valid {
			dataRowProduct.Stock = dataRowProductStock.Int64
		} else {
			dataRowProduct.Stock = 0
		}
		result = append(result, dataRowProduct)
	}
	return result, nil
}

func (pr *ProductRepo) AddProduct(newUser *model.Product) (*model.Product, error) {
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

	return &model.Product{
		ID:    int(lastInsertedID),
		Name:  newUser.Name,
		Price: newUser.Price,
	}, nil
}

func (pr *ProductRepo) AddProductStock(addQuantity *model.ProductIncrease) (*model.ProductIncrease, error) {
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
		conn.Prepare("insert into product_increases(product_id, quantity, user_id) values (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(addQuantity.ProductID, addQuantity.Quantity, addQuantity.UserID)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.ProductIncrease{
		ID:        int(lastInsertedID),
		ProductID: addQuantity.ProductID,
		Quantity:  addQuantity.Quantity,
		UserID:    addQuantity.UserID,
	}, nil
}

func (pr *ProductRepo) DecProductStock(decQuantity *model.ProductDec) (*model.ProductDec, error) {
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

	return &model.ProductDec{
		ID:        int(lastInsertedID),
		ProductID: decQuantity.ProductID,
		Quantity:  decQuantity.Quantity,
		InvoiceID: decQuantity.InvoiceID,
	}, nil
}
