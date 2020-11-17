package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
)

type ProductCTR struct {
	ProductUseCase usecase.ProductUseCase
	Res            customJSON.JSONCustom
}

func (pctr *ProductCTR) GetProducts(w http.ResponseWriter, _ *http.Request) {
	result, err := pctr.ProductUseCase.GetProducts()
	if err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		200, "success", "", result)
	return
}

func (pctr *ProductCTR) AddProduct(w http.ResponseWriter, r *http.Request) {
	var newUser model.Product
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusBadRequest, "error", err.Error(), nil)
		return
	}
	result, err := pctr.ProductUseCase.AddProduct(&newUser)
	if err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}
	pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		200, "success", "", result)
	return
}
