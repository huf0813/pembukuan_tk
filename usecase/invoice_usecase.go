package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type InvoiceUseCase struct {
	InvoiceRepo    sqlite.InvoiceRepo
	ProductUseCase ProductUseCase
}

func (iuc *InvoiceUseCase) InvoiceValidation(newInvoiceReq *model.InvoiceReq) error {
	err := errors.New("cannot pass the validation")
	if newInvoiceReq.ListProduct == nil {
		return err
	}
	if newInvoiceReq.UserID <= 0 {
		return err
	}
	if newInvoiceReq.CustomerID <= 0 {
		return err
	}
	return nil
}

func (iuc *InvoiceUseCase) ProductDecValidation(newProdDec *model.ProductDec) error {
	if newProdDec.Quantity <= 0 {
		return errors.New("quantity cannot be less than equal zero")
	}
	if newProdDec.ProductID <= 0 {
		return errors.New("product is not available")
	}
	if newProdDec.InvoiceID <= 0 {
		return errors.New("invoice is not available")
	}
	getProducts, err := iuc.ProductUseCase.GetProducts()
	if err != nil {
		return err
	}
	for _, v := range getProducts {
		if newProdDec.ProductID == v.ID {
			if (int(v.Stock) - newProdDec.Quantity) < 0 {
				return errors.New("stock is not available")
			}
		}
	}
	return nil
}

func (iuc *InvoiceUseCase) AddInvoice(newInvoiceReq *model.InvoiceReq) (*model.Invoice, error) {
	if err := iuc.InvoiceValidation(newInvoiceReq); err != nil {
		return nil, err
	}
	newInvoice := &model.Invoice{
		UserID:     newInvoiceReq.UserID,
		CustomerID: newInvoiceReq.CustomerID,
	}

	resultInvoice, err := iuc.InvoiceRepo.AddInvoice(newInvoice)
	if err != nil {
		return nil, err
	}

	for _, val := range newInvoiceReq.ListProduct {
		obj := &model.ProductDec{
			ProductID: val.ProductID,
			Quantity:  val.Qty,
			InvoiceID: resultInvoice.ID,
		}
		if err := iuc.ProductDecValidation(obj); err != nil {
			return nil, err
		}
		if _, err := iuc.ProductUseCase.DecProductStock(obj); err != nil {
			return nil, err
		}
	}

	return &model.Invoice{
		ID:         resultInvoice.ID,
		CustomerID: newInvoiceReq.CustomerID,
		UserID:     newInvoiceReq.UserID,
	}, nil
}
