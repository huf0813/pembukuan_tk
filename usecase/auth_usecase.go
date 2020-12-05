package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/middleware"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
	"github.com/huf0813/pembukuan_tk/utils"
)

type AuthUseCase struct {
	Auth     middleware.TokenMiddleware
	UserRepo sqlite.UserRepo
	Hashing  utils.Hashing
}

type AuthUseCaseInterface interface {
	LoginValidation(user *entity.UserReq) error
	Login(userAuth *entity.UserReq) (*entity.TokenExtract, error)
	CheckTokenExpired(token string) error
}

func (auc *AuthUseCase) LoginValidation(user *entity.UserReq) error {
	if user.Username == "" || user.Password == "" {
		return errors.New("field is empty")
	}
	return nil
}

func (auc *AuthUseCase) Login(userAuth *entity.UserReq) (*entity.TokenExtract, error) {
	if err := auc.LoginValidation(userAuth); err != nil {
		return nil, err
	}

	userFoundInDatabase, err := auc.UserRepo.FindUserByUsername(userAuth.Username)
	if err != nil {
		return nil, err
	}
	if err := auc.Hashing.ComparePass(userFoundInDatabase.Password, userAuth.Password); err != nil {
		return nil, err
	}

	result, err := auc.Auth.GetToken(userFoundInDatabase.Username, userFoundInDatabase.UserTypeID, userFoundInDatabase.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (auc *AuthUseCase) CheckTokenExpired(token string) error {
	if _, err := auc.Auth.VerifyToken(token); err != nil {
		return err
	}
	return nil
}
