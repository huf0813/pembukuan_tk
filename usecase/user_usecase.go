package usecase

import (
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
	"github.com/huf0813/pembukuan_tk/utils"
)

type UserUseCase struct {
	UserRepo sqlite.UserRepo
	Hash     utils.Hashing
}

type UserUseCaseInterface interface {
	GetUsers() ([]model.User, error)
	AddUser() (*model.User, error)
}

func (uuc *UserUseCase) GetUsers() ([]model.User, error) {
	return uuc.UserRepo.GetUsers()
}

func (uuc *UserUseCase) AddUser(username, password string, userTypeID int) (*model.User, error) {
	hashPass, err := uuc.Hash.HashPass(password)
	if err != nil {
		return nil, err
	}
	newUser := &model.User{
		Username:   username,
		Password:   string(hashPass),
		UserTypeID: userTypeID,
	}
	result, err := uuc.UserRepo.AddUser(newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}
