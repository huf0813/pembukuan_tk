package ctr

import (
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/utils/delivery"
	"net/http"
)

type HomeCTR struct {
	Res     delivery.CustomJSONUtil
	Payload model.Payload
}

type HomeCTRInterface interface {
	Welcome(w http.ResponseWriter, r *http.Request)
}

func (h *HomeCTR) Welcome(w http.ResponseWriter, _ *http.Request) {
	h.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, h.Payload.NewPayload("success", "", "welcome to api!"))
}
