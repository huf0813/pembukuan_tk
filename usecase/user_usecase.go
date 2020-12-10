package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
	"github.com/huf0813/pembukuan_tk/utils"
)

type UserUseCase struct {
	UserRepo sqlite.UserRepo
	Hash     utils.Hashing
}

type UserUseCaseInterface interface {
	GetUsers() ([]entity.User, error)
	AddUserValidation(newUser *entity.User) error
	AddUser(username, password string, userTypeID int) (*entity.User, error)
	EditUserValidation(newUser *entity.User) error
	EditUser(editUsername, editedPassword string, userID int) (*entity.User, error)
	GetAllUsersIncludingAdmin() ([]entity.User, error)
}

func (uuc *UserUseCase) GetUsers() ([]entity.User, error) {
	return uuc.UserRepo.GetUsers()
}

func (uuc *UserUseCase) GetAllUsersIncludingAdmin() ([]entity.User, error) {
	return uuc.UserRepo.GetAllUsers()
}

func (uuc *UserUseCase) AddUserValidation(newUser *entity.User) error {
	if newUser.Password == "" {
		return errors.New("password cannot be empty")
	}
	if newUser.Username == "" {
		return errors.New("username cannot be empty")
	}
	return nil
}

func (uuc *UserUseCase) AddUser(username, password string, userTypeID int) (*entity.User, error) {
	hashPass, err := uuc.Hash.HashPass(password)
	if err != nil {
		return nil, err
	}
	newUser := &entity.User{
		Username:   username,
		Password:   string(hashPass),
		UserTypeID: userTypeID,
	}
	if err := uuc.AddUserValidation(newUser); err != nil {
		return nil, err
	}
	result, err := uuc.UserRepo.AddUser(newUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uuc *UserUseCase) EditUserValidation(newUser *entity.User) error {
	if newUser.Password == "" {
		return errors.New("password cannot be empty")
	}
	if newUser.Username == "" {
		return errors.New("username cannot be empty")
	}
	if newUser.ID == 0 {
		return errors.New("username cannot be empty")
	}

	flag := false
	res, err := uuc.GetUsers()
	if err != nil {
		return err
	}
	for _, val := range res {
		if val.ID == newUser.ID {
			flag = true
		}
	}
	if !flag {
		return errors.New("user was not found")
	}
	return nil
}

func (uuc *UserUseCase) EditUser(editUsername, editedPassword string, userID int) (*entity.User, error) {
	hashPass, err := uuc.Hash.HashPass(editedPassword)
	if err != nil {
		return nil, err
	}
	editedUser := &entity.User{
		ID:         userID,
		Username:   editUsername,
		Password:   string(hashPass),
		UserTypeID: 2,
	}
	if err := uuc.EditUserValidation(editedUser); err != nil {
		return nil, err
	}
	result, err := uuc.UserRepo.EditUser(editedUser)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (uuc *UserUseCase) DeleteUser(userID int) (string, error) {
	result, err := uuc.UserRepo.DeleteUser(userID)
	if err != nil {
		return "", err
	}
	return result, nil
}
