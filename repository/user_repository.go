package repository

import "github.com/huf0813/pembukuan_tk/model"

type UserRepo struct{}

type UserRepoInterface interface {
	GetUsers() ([]model.User, error)
}

func (ur *UserRepo) GetUsers() ([]model.User, error) {
	return []model.User{
		{"Harun", "huf0813", "harun@gmail.com", "08123456789", "masih dummy"},
		{"Joseph", "joseph88888", "joseph@gmail.com", "08123456789", "masih dummy"},
	}, nil
}
