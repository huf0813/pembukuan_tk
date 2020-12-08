package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/utils"
	"time"
)

type UserRepo struct {
	Timestamp utils.Timestamp
	SqlConn   sqlite.ConnSqlite
}

type UserRepoInterface interface {
	GetUsers() ([]entity.User, error)
	FindUserByUsername(usernameSearch string) (*entity.User, error)
	AddUser(user *entity.User) (*entity.User, error)
	EditUser(editedUser *entity.User) (*entity.User, error)
}

func (ur *UserRepo) GetUsers() ([]entity.User, error) {
	conn := ur.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	rows, err := conn.Query("select id, user_type_id, username, password from users where user_type_id=2")
	if err != nil {
		return nil, err
	}

	var result []entity.User
	for rows.Next() {
		var rowData entity.User
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

func (ur *UserRepo) FindUserByUsername(usernameSearch string) (*entity.User, error) {
	conn := ur.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	var founded entity.User
	if err := conn.QueryRow("select id, user_type_id, username, password from users where username=?",
		usernameSearch).Scan(&founded.ID, &founded.UserTypeID, &founded.Username, &founded.Password); err != nil {
		return nil, err
	}
	return &founded, nil
}

func (ur *UserRepo) AddUser(newUser *entity.User) (*entity.User, error) {
	conn := ur.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("insert into users(user_type_id, username, password, created_at, updated_at) values (?, ?, ?, ?, ?)")
	if err != nil {
		return nil, err
	}
	getID, err := result.Exec(newUser.UserTypeID, newUser.Username, newUser.Password, time.Now().Unix(), time.Now().Unix())
	if err != nil {
		return nil, err
	}
	lastInsertedID, err := getID.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &entity.User{
		ID:         int(lastInsertedID),
		Username:   newUser.Username,
		Password:   newUser.Password,
		UserTypeID: newUser.UserTypeID,
	}, nil
}

func (ur *UserRepo) EditUser(editedUser *entity.User) (*entity.User, error) {
	conn := ur.SqlConn.SqliteConnInit()
	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	if conn == nil {
		return nil, errors.New("connection failed to db")
	}

	result, err :=
		conn.Prepare("update users set username=?, password=?, updated_at=? where id=?")
	if err != nil {
		return nil, err
	}
	edited, err := result.Exec(editedUser.Username, editedUser.Password, time.Now().Unix(), editedUser.ID)
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

	return &entity.User{
		ID:         editedUser.ID,
		Username:   editedUser.Username,
		Password:   editedUser.Password,
		UserTypeID: editedUser.UserTypeID,
	}, nil
}
