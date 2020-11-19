package usecase

import (
	"github.com/huf0813/pembukuan_tk/model"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type ProductUseCase struct {
	ProductRepo sqlite.ProductRepo
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
	result, err := pus.ProductRepo.AddProductStock(addProductStock)
	if err != nil {
		return nil, err
	}
	return result, err
}
