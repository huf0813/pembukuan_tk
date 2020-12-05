package ctr

import (
	"encoding/json"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/usecase"
	"github.com/huf0813/pembukuan_tk/utils/delivery/customJSON"
	"net/http"
)

type InvoiceCTR struct {
	InvoiceUseCase usecase.InvoiceUseCase
	Res            customJSON.JSONCustom
}

func (ictr *InvoiceCTR) AddNewInvoice(w http.ResponseWriter, r *http.Request) {
	var invoiceReq entity.InvoiceReq
	if err := json.NewDecoder(r.Body).Decode(&invoiceReq); err != nil {
		ictr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusBadRequest, "error", err.Error(), nil)
		return
	}

	result, err := ictr.InvoiceUseCase.AddInvoice(&invoiceReq)
	if err != nil {
		ictr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	ictr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success", "", result)
	return
}

func (ictr *InvoiceCTR) GetInvoices(w http.ResponseWriter, _ *http.Request) {
	result, err := ictr.InvoiceUseCase.GetInvoices()
	if err != nil {
		ictr.Res.CustomJSONRes(w, "Content-Type", "application/json",
			http.StatusInternalServerError, "error", err.Error(), nil)
		return
	}

	ictr.Res.CustomJSONRes(w, "Content-Type", "application/json",
		http.StatusOK, "success", "", result)
	return
}
