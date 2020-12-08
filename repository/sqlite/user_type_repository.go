package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/utils"
	"time"
)

type UserTypeRepo struct {
	Timestamp utils.Timestamp
	SqlConn   sqlite.ConnSqlite
}

func (utr *UserTypeRepo) AddUserType(newUserType string) error {
	conn := utr.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into user_types(name, created_at, updated_at) values (?,?,?)")
	if err != nil {
		return err
	}
	if _, err := result.Exec(newUserType, time.Now().Unix(), time.Now().Unix()); err != nil {
		return err
	}
	return nil
}
