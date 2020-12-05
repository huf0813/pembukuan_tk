package ctr

import (
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
)

type AdminCTR struct {
	Res customJSON.JSONCustom
}

type AdminCTRInterface interface {
	DashboardAdmin(w http.ResponseWriter, r *http.Request)
}

func (ac *AdminCTR) DashboardAdmin(w http.ResponseWriter, _ *http.Request) {
	ac.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "welcome to admin dashboard", "")
}
