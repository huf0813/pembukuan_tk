package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
)

type AuthCTR struct {
	Res         customJSON.JSONCustom
	AuthUseCase usecase.AuthUseCase
}

type AuthInterface interface {
	Login(w http.ResponseWriter, r *http.Request)
}

func (ac *AuthCTR) Login(w http.ResponseWriter, r *http.Request) {
	var authLogin entity.UserReq
	if err := json.NewDecoder(r.Body).Decode(&authLogin); err != nil {
		ac.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}

	result, err := ac.AuthUseCase.Login(&authLogin)
	if err != nil {
		ac.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	ac.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", result)
	return
}

func (ac *AuthCTR) CheckToken(w http.ResponseWriter, r *http.Request) {
	var tokenReq entity.TokenReq
	if err := json.NewDecoder(r.Body).Decode(&tokenReq); err != nil {
		ac.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}

	if err := ac.AuthUseCase.CheckTokenExpired(tokenReq.Token); err != nil {
		ac.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	ac.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "token is valid", nil)
	return
}
