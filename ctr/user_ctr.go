package ctr

import (
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
}

func (uc *UserCTR) DashboardUser(w http.ResponseWriter, _ *http.Request) {
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", "welcome to users dashboard")
	return
}
