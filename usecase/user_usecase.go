package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/middleware"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository"
)

type UserUseCase struct {
	Auth     middleware.TokenMiddleware
	UserRepo repository.UserRepo
}

type UserUseCaseInterface interface {
	ValidateUser(user *model.User) error
	Login(user *model.User) (string, error)
	GetUsers() ([]model.User, error)
}

func (uuc *UserUseCase) ValidateUser(user *model.UserLogin) error {
	if user.Username == "" {
		return errors.New("username is empty")
	}

	if user.Password == "" {
		return errors.New("password is empty")
	}

	return nil
}

func (uuc *UserUseCase) Login(user *model.UserLogin) (string, error) {
	if err := uuc.ValidateUser(user); err != nil {
		return "", err
	}

	result, err := uuc.Auth.GetToken(user.Username)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (uuc *UserUseCase) GetUsers() ([]model.User, error) {
	return uuc.UserRepo.GetUsers()
}
