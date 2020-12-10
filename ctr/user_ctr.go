package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
)

type UserCTR struct {
	Res         customJSON.JSONCustom
	UserUseCase usecase.UserUseCase
}

type UserCTRInterface interface {
	DashboardUser(w http.ResponseWriter, r *http.Request)
	FetchUsers(w http.ResponseWriter, r *http.Request)
	AddUser(w http.ResponseWriter, r *http.Request)
	EditedUser(w http.ResponseWriter, r *http.Request)
	DeletedUser(w http.ResponseWriter, r *http.Request)
}

func (uc *UserCTR) DashboardUser(w http.ResponseWriter, _ *http.Request) {
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "welcome to users dashboard", nil)
	return
}

func (uc *UserCTR) FetchUsers(w http.ResponseWriter, _ *http.Request) {
	result, err := uc.UserUseCase.GetUsers()
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success", "", result)
	return
}

func (uc *UserCTR) AddUser(w http.ResponseWriter, r *http.Request) {
	var newUser entity.UserReq
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}

	result, err := uc.UserUseCase.AddUser(newUser.Username, newUser.Password, 2)
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success", "", result)
	return
}

func (uc *UserCTR) EditedUser(w http.ResponseWriter, r *http.Request) {
	var editedUser entity.User
	if err := json.NewDecoder(r.Body).Decode(&editedUser); err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}

	result, err := uc.UserUseCase.EditUser(editedUser.Username, editedUser.Password, editedUser.ID)
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success", "", result)
	return
}

func (uc *UserCTR) DeletedUser(w http.ResponseWriter, r *http.Request) {
	var editedUser entity.DeleteRowTemp
	if err := json.NewDecoder(r.Body).Decode(&editedUser); err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}

	result, err := uc.UserUseCase.DeleteUser(editedUser.ID)
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success", result, nil)
	return
}
