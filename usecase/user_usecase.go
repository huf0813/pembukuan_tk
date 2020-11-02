package usecase

import (
	"github.com/huf0813/pembukuan_tk/middleware"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type UserUseCase struct {
	Auth     middleware.TokenMiddleware
	UserRepo sqlite.UserRepo
}

type UserUseCaseInterface interface {
	GetUsers() (*[]model.User, error)
}

func (uuc *UserUseCase) GetUsers() (*[]model.User, error) {
	return uuc.UserRepo.GetUsers()
}
