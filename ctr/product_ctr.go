package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/entity"
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
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		200, "success", "", result)
	return
}

func (pctr *ProductCTR) AddProduct(w http.ResponseWriter, r *http.Request) {
	var newProduct entity.Product
	if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}

	result, err := pctr.ProductUseCase.AddProduct(&newProduct)
	if err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		200, "success", "", result)
	return
}

func (pctr *ProductCTR) EditProduct(w http.ResponseWriter, r *http.Request) {
	var editedProduct entity.Product
	if err := json.NewDecoder(r.Body).Decode(&editedProduct); err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}

	result, err := pctr.ProductUseCase.EditProduct(&editedProduct)
	if err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		200, "success", "", result)
	return
}

func (pctr *ProductCTR) AddProductStock(w http.ResponseWriter, r *http.Request) {
	var addProductStock *entity.ProductIncrease
	if err := json.NewDecoder(r.Body).Decode(&addProductStock); err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}

	result, err := pctr.ProductUseCase.AddProductStock(addProductStock)
	if err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		200, "success", "", result)
	return
}

func (pctr *ProductCTR) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	var deleteProduct *entity.DeleteRowTemp
	if err := json.NewDecoder(r.Body).Decode(&deleteProduct); err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}

	res, err := pctr.ProductUseCase.DeleteProduct(deleteProduct.ID)
	if err != nil {
		pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusOK, "error", err.Error(), nil)
		return
	}
	pctr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		200, "success", "", res)
	return
}
