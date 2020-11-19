package sqlite

import (
	"database/sql"
	"errors"
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
	rows, err := conn.Query("select id, name, price, (select name from product_types WHERE product_types.id = products.product_type_id) product_type, (select sum(product_increases.quantity) from product_increases WHERE product_increases.product_id = products.id) stock from products")
	if err != nil {
		return nil, err
	}

	var result []model.ProductStockAndType
	for rows.Next() {
		var dataRowProduct model.ProductStockAndType
		var dataRowProduct_Stock sql.NullInt64
		if err := rows.Scan(&dataRowProduct.ID,
			&dataRowProduct.Name,
			&dataRowProduct.Price,
			&dataRowProduct.ProductType,
			&dataRowProduct_Stock); err != nil {
			return nil, err
		}
		if dataRowProduct_Stock.Valid {
			dataRowProduct.Stock = dataRowProduct_Stock.Int64
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
