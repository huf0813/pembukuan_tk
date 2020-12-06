package usecase

import (
	"errors"
	"fmt"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
	"strconv"
	"strings"
	"time"
)

type InvoiceUseCase struct {
	InvoiceRepo    sqlite.InvoiceRepo
	ProductUseCase ProductUseCase
}

func (iuc *InvoiceUseCase) GetInvoices() ([]entity.InvoiceWithDetail, error) {
	result, err := iuc.InvoiceRepo.GetInvoices()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (iuc *InvoiceUseCase) GetInvoiceByID(invoiceID int) (*entity.InvoiceWithDetail, error) {
	result, err := iuc.InvoiceRepo.GetInvoices()
	if err != nil {
		return nil, err
	}
	for _, val := range result {
		if val.ID == invoiceID {
			return &val, nil
		}
	}
	return nil, errors.New("invoice was not found")
}

func (iuc *InvoiceUseCase) InvoiceValidation(newInvoiceReq *entity.InvoiceReq) error {
	err := errors.New("cannot pass the validation")
	if newInvoiceReq.ListProduct == nil {
		return err
	}
	if len(newInvoiceReq.ListProduct) == 0 {
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

func (iuc *InvoiceUseCase) InvoiceProductDecValidation(newProdDec *entity.ProductDec) error {
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

func (iuc *InvoiceUseCase) AddInvoice(newInvoiceReq *entity.InvoiceReq) (*entity.Invoice, error) {
	if err := iuc.InvoiceValidation(newInvoiceReq); err != nil {
		return nil, err
	}
	newInvoice := &entity.Invoice{
		UserID:     newInvoiceReq.UserID,
		CustomerID: newInvoiceReq.CustomerID,
	}

	resultInvoice, err := iuc.InvoiceRepo.AddInvoice(newInvoice)
	if err != nil {
		return nil, err
	}

	for _, val := range newInvoiceReq.ListProduct {
		obj := &entity.ProductDec{
			ProductID: val.ProductID,
			Quantity:  val.Qty,
			InvoiceID: resultInvoice.ID,
		}
		if err := iuc.InvoiceProductDecValidation(obj); err != nil {
			return nil, err
		}
		if _, err := iuc.ProductUseCase.DecProductStock(obj); err != nil {
			return nil, err
		}
	}

	return &entity.Invoice{
		ID:         resultInvoice.ID,
		CustomerID: newInvoiceReq.CustomerID,
		UserID:     newInvoiceReq.UserID,
	}, nil
}

func (iuc *InvoiceUseCase) GetStatistics(year string) (*entity.StatisticPerMonRes, error) {
	result, err := iuc.InvoiceRepo.GetInvoices()
	if err != nil {
		return nil, err
	}
	profitPerMonth := []entity.StatisticPerYear{
		{fmt.Sprintf("%s-01", year), 0},
		{fmt.Sprintf("%s-02", year), 0},
		{fmt.Sprintf("%s-03", year), 0},
		{fmt.Sprintf("%s-04", year), 0},
		{fmt.Sprintf("%s-05", year), 0},
		{fmt.Sprintf("%s-06", year), 0},
		{fmt.Sprintf("%s-07", year), 0},
		{fmt.Sprintf("%s-08", year), 0},
		{fmt.Sprintf("%s-09", year), 0},
		{fmt.Sprintf("%s-10", year), 0},
		{fmt.Sprintf("%s-11", year), 0},
		{fmt.Sprintf("%s-12", year), 0},
	}
	for _, val := range result {
		convertInt64, err := iuc.StringToInt64(val.CreatedAt)
		if err != nil {
			return nil, err
		}
		convertYnM, err := iuc.GetYnM(convertInt64)
		if err != nil {
			return nil, err
		}
		for i, perM := range profitPerMonth {
			res, err := iuc.StringToInt64(val.TotalInvoicePrice)
			if err != nil {
				return nil, err
			}
			if convertYnM == perM.YearAndMon {
				profitPerMonth[i].Profit += res
			}
		}
	}
	return &entity.StatisticPerMonRes{
		Year: year,
		Detail: []entity.StatisticPerMon{
			{"january", profitPerMonth[0].Profit},
			{"february", profitPerMonth[1].Profit},
			{"march", profitPerMonth[2].Profit},
			{"april", profitPerMonth[3].Profit},
			{"may", profitPerMonth[4].Profit},
			{"june", profitPerMonth[5].Profit},
			{"july", profitPerMonth[6].Profit},
			{"august", profitPerMonth[7].Profit},
			{"september", profitPerMonth[8].Profit},
			{"october", profitPerMonth[9].Profit},
			{"november", profitPerMonth[10].Profit},
			{"december", profitPerMonth[11].Profit},
		},
	}, nil
}

func (iuc *InvoiceUseCase) StringToInt64(val string) (int64, error) {
	if n, err := strconv.ParseInt(val, 10, 64); err == nil {
		return n, nil
	} else {
		return 0, err
	}
}

func (iuc *InvoiceUseCase) GetYnM(timeUnix int64) (string, error) {
	timeT := time.Unix(timeUnix, 0).String()
	getYMD := strings.Split(timeT, " ")[0]
	getYM := strings.Split(getYMD, "-")[:2]
	stringYM := fmt.Sprintf("%s-%s", getYM[0], getYM[1])
	return stringYM, nil
}
