package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/middleware"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type UserUseCase struct {
	Auth     middleware.TokenMiddleware
	UserRepo sqlite.UserRepo
}

type UserUseCaseInterface interface {
	ValidateUser(user *model.UserLogin) error
	Login(userAuth *model.UserLogin) (string, error)
	GetUsers() (*[]model.User, error)
}

func (uuc *UserUseCase) ValidateUser(user *model.UserLogin) error {
	if user.Username == "" && user.Password == "" {
		return errors.New("fields have to be filled")
	}
	return nil
}

func (uuc *UserUseCase) Login(userAuth *model.UserLogin) (string, error) {
	if err := uuc.ValidateUser(userAuth); err != nil {
		return "", err
	}

	userFoundInDatabase, err := uuc.UserRepo.FindUserByUsername(userAuth.Username)
	if err != nil {
		return "", err
	}
	if userFoundInDatabase.Password != userAuth.Password {
		return "", errors.New("username or password is wrong")
	}

	result, err := uuc.Auth.GetToken(userAuth.Username)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (uuc *UserUseCase) GetUsers() (*[]model.User, error) {
	return uuc.UserRepo.GetUsers()
}
