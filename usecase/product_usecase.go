package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type ProductUseCase struct {
	ProductRepo sqlite.ProductRepo
}

func (pus *ProductUseCase) AddProductValidation(productInc *model.ProductIncrease) error {
	var err = errors.New("can't pass the validation")
	if productInc.Quantity < 0 {
		return err
	}
	return nil
}

func (pus *ProductUseCase) DecProductValidation(productDec *model.ProductDec) error {
	var err = errors.New("can't pass the validation")
	if productDec.Quantity < 0 {
		return err
	}
	return nil
}

func (pus *ProductUseCase) GetProducts() ([]model.ProductStockAndType, error) {
	result, err := pus.ProductRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pus *ProductUseCase) AddProduct(newProduct *model.Product) (*model.Product, error) {
	result, err := pus.ProductRepo.AddProduct(newProduct)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pus *ProductUseCase) AddProductStock(addProductStock *model.ProductIncrease) (*model.ProductIncrease, error) {
	if err := pus.AddProductValidation(addProductStock); err != nil {
		return nil, err
	}

	result, err := pus.ProductRepo.AddProductStock(addProductStock)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (pus *ProductUseCase) DecProductStock(decProductStock *model.ProductDec) (*model.ProductDec, error) {
	if err := pus.DecProductValidation(decProductStock); err != nil {
		return nil, err
	}

	result, err := pus.ProductRepo.DecProductStock(decProductStock)
	if err != nil {
		return nil, err
	}

	return result, err
}
