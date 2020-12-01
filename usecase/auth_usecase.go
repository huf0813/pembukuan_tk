package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/middleware"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type AuthUseCase struct {
	Auth     middleware.TokenMiddleware
	UserRepo sqlite.UserRepo
}

type AuthUseCaseInterface interface {
	ValidateUser(user *model.AuthLogin) error
	Login(userAuth *model.AuthLogin) (*model.Token, error)
}

func (auc *AuthUseCase) ValidateUser(user *model.AuthLogin) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("field is empty")
	}
	return nil
}

func (auc *AuthUseCase) Login(userAuth *model.AuthLogin) (*model.TokenExtract, error) {
	if err := auc.ValidateUser(userAuth); err != nil {
		return nil, err
	}

	userFoundInDatabase, err := auc.UserRepo.FindUserByUsername(userAuth.Username)
	if err != nil {
		return nil, err
	}
	if userFoundInDatabase.Password != userAuth.Password {
		return nil, errors.New("username or password is wrong")
	}

	result, err := auc.Auth.GetToken(userFoundInDatabase.Username, userFoundInDatabase.UserTypeID, userFoundInDatabase.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
