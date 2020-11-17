package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/model"
)

type ProductRepo struct {
	SqlConn sqlite.ConnSqlite
}

func (pr *ProductRepo) GetProducts() ([]model.Product, error) {
	conn := pr.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}
	rows, err := conn.Query("select  * from products")
	if err != nil {
		return nil, err
	}

	var result []model.Product
	for rows.Next() {
		var dataRowProduct model.Product
		if err := rows.Scan(&dataRowProduct.ID, &dataRowProduct.Name, &dataRowProduct.Price, &dataRowProduct.ProductTypeID); err != nil {
			return nil, err
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
		conn.Prepare("insert into products(name, price, product_type_id) values (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(newUser.Name, newUser.Price, newUser.ProductTypeID)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.Product{
		ID:            int(lastInsertedID),
		Name:          newUser.Name,
		Price:         newUser.Price,
		ProductTypeID: newUser.ProductTypeID,
	}, nil
}
