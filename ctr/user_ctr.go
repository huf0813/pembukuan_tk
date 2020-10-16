package ctr

import (
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery"
	"net/http"
)

type UserCTR struct {
	Res         delivery.CustomJSONUtil
	Payload     model.Payload
	UserUseCase usecase.UserUseCase
}

type UserCTRInterface interface {
	GetUsers(w http.ResponseWriter, r *http.Request)
}

func (uc *UserCTR) GetUsers(w http.ResponseWriter, _ *http.Request) {
	result, err := uc.UserUseCase.FetchWithoutID()
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusInternalServerError, uc.Payload.NewPayload("error", err.Error(), nil))
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, uc.Payload.NewPayload("success", "", result))
	return
}
