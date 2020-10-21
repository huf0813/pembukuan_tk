package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery"
	"net/http"
)

type UserCTR struct {
	Res         delivery.CustomJSONUtil
	UserUseCase usecase.UserUseCase
}

type UserCTRInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
	GetUsers(w http.ResponseWriter, r *http.Request)
	TestingGetUsers(w http.ResponseWriter, r *http.Request)
}

func (uc *UserCTR) Login(w http.ResponseWriter, r *http.Request) {
	var userLogin model.UserLogin
	if err := json.NewDecoder(r.Body).Decode(&userLogin); err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	result, err := uc.UserUseCase.Login(&userLogin)
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", result)
	return
}

func (uc *UserCTR) GetUsers(w http.ResponseWriter, _ *http.Request) {
	result, err := uc.UserUseCase.GetUsers()
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", result)
	return
}

func (uc *UserCTR) TestingGetUsers(w http.ResponseWriter, _ *http.Request) {
	result, err := uc.UserUseCase.GetUsers()
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", result)
	return
}
