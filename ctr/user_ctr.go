package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/model"
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
	CustomerRegister(w http.ResponseWriter, r *http.Request)
	FetchCustomers(w http.ResponseWriter, r *http.Request)
}

func (uc *UserCTR) DashboardUser(w http.ResponseWriter, _ *http.Request) {
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", "welcome to users dashboard")
	return
}

func (uc *UserCTR) CustomerRegister(w http.ResponseWriter, r *http.Request) {
	var newCustomer model.Customer
	if err := json.NewDecoder(r.Body).Decode(&newCustomer); err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	insertedCustomer, err := uc.CustomerUseCase.AddNewCustomer(newCustomer.Name, newCustomer.Phone, newCustomer.Email, newCustomer.Address)
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "inserted successfully", insertedCustomer)
	return
}

func (uc *UserCTR) FetchCustomers(w http.ResponseWriter, _ *http.Request) {
	result, err := uc.CustomerUseCase.GetCustomers()
	if err != nil {
		uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	uc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", result)
	return
}
