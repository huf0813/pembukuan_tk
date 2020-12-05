package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
	"time"
)

type ProductIncreaseRepo struct {
	SqlConn sqlite.ConnSqlite
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
