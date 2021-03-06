package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
)

type CustomerCTR struct {
	CustomerUseCase usecase.CustomerUseCase
	Res             customJSON.JSONCustom
}

func (cc *CustomerCTR) FetchCustomers(w http.ResponseWriter, _ *http.Request) {
	result, err := cc.CustomerUseCase.GetCustomers()
	if err != nil {
		cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "", result)
	return
}

func (cc *CustomerCTR) CustomerRegister(w http.ResponseWriter, r *http.Request) {
	var newCustomer entity.Customer
	if err := json.NewDecoder(r.Body).Decode(&newCustomer); err != nil {
		cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}

	insertedCustomer, err := cc.CustomerUseCase.AddCustomer(newCustomer.Name,
		newCustomer.Phone,
		newCustomer.Email,
		newCustomer.Address)
	if err != nil {
		cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}

	cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "success", "inserted successfully", insertedCustomer)
	return
}

func (cc *CustomerCTR) EditCustomer(w http.ResponseWriter, r *http.Request) {
	var editCustomer entity.Customer
	if err := json.NewDecoder(r.Body).Decode(&editCustomer); err != nil {
		cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}

	updatedCustomer, err := cc.CustomerUseCase.EditCustomer(editCustomer.Name,
		editCustomer.Phone,
		editCustomer.Email,
		editCustomer.Address,
		editCustomer.ID)
	if err != nil {
		cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}
	cc.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success",
		"edited successfully",
		updatedCustomer)
	return
}

func (cc *CustomerCTR) DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	var editCustomer entity.DeleteRowTemp
	if err := json.NewDecoder(r.Body).Decode(&editCustomer); err != nil {
		cc.Res.CustomJSONRes(w, "Content-Type", "application/json", http.StatusOK, "error", err.Error(), nil)
		return
	}

	res, err := cc.CustomerUseCase.DeleteCustomer(editCustomer.ID)
	if err != nil {
		cc.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	cc.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success",
		res,
		nil)
	return
}
