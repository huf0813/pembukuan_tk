package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/utils"
)

type UserRepo struct {
	Timestamp utils.Timestamp
	SqlConn   sqlite.ConnSqlite
}

type UserRepoInterface interface {
	GetUsers() ([]model.User, error)
	FindUserByUsername(usernameSearch string) (*model.User, error)
	AddUser(user *model.User) (*model.User, error)
}

func (ur *UserRepo) GetUsers() ([]model.User, error) {
	conn := ur.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	rows, err := conn.Query("select id, user_type_id, username, password from users")
	if err != nil {
		return nil, err
	}

	var result []model.User
	for rows.Next() {
		var rowData model.User
		if err := rows.Scan(&rowData.ID,
			&rowData.UserTypeID,
			&rowData.Username,
			&rowData.Password); err != nil {
			return nil, err
		}
		result = append(result, rowData)
	}

	return result, nil
}

func (ur *UserRepo) FindUserByUsername(usernameSearch string) (*model.User, error) {
	conn := ur.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	var founded model.User
	if err := conn.QueryRow("select id, user_type_id, username, password from users where username=?",
		usernameSearch).Scan(&founded.ID, &founded.UserTypeID, &founded.Username, &founded.Password); err != nil {
		return nil, err
	}
	return &founded, nil
}

func (ur *UserRepo) AddUser(newUser *model.User) (*model.User, error) {
	conn := ur.SqlConn.SqliteConn()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into users(user_type_id, username, password) values (?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(newUser.UserTypeID, newUser.Username, newUser.Password)
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:         int(lastInsertedID),
		Username:   newUser.Username,
		Password:   newUser.Password,
		UserTypeID: newUser.UserTypeID,
	}, nil
}
