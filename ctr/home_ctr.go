package ctr

import (
	"github.com/huf0813/pembukuan_tk/utils"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
	"time"
)

type HomeCTR struct {
	Res       customJSON.JSONCustom
	Timestamp utils.Timestamp
}

type HomeCTRInterface interface {
	Welcome(w http.ResponseWriter, _ *http.Request)
}

func (h *HomeCTR) Welcome(w http.ResponseWriter, _ *http.Request) {
	h.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", "Welcome to API at "+h.Timestamp.PrettyTime(time.Now()))
}
