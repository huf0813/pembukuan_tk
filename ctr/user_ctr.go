package ctr

import (
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
)

type UserCTR struct {
	Res             customJSON.JSONCustom
	UserUseCase     usecase.UserUseCase
	CustomerUseCase usecase.CustomerUseCase
}

type UserCTRInterface interface {
	DashboardUser(w http.ResponseWriter, r *http.Request)
	FetchUsers(w http.ResponseWriter, r *http.Request)
}

func (uc *UserCTR) DashboardUser(w http.ResponseWriter, _ *http.Request) {
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", "welcome to users dashboard")
	return
}

func (uc *UserCTR) FetchUsers(w http.ResponseWriter, _ *http.Request) {
	result, err := uc.UserUseCase.GetUsers()
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success", "", result)
	return
}
