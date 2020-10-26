package sqlite

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/utils"
)

type UserRepo struct {
	Timestamp utils.Timestamp
}

type UserRepoInterface interface {
	GetUsers() (*[]model.User, error)
	FindUserByUsername() (*model.User, bool)
}

var users = []model.User{
	{"Haroun", "huf0813", "harun@gmail.com", "08123456789", "dummy"},
	{"Joseph", "joseph88888", "joseph@gmail.com", "08123456789", "dummy"},
}

func (ur *UserRepo) GetUsers() (*[]model.User, error) {
	return &users, nil
}

func (ur *UserRepo) FindUserByUsername(usernameSearch string) (*model.User, error) {
	for _, user := range users {
		if user.Username == usernameSearch {
			return &user, nil
		}
	}
	return nil, errors.New("user is not registered")
}
