package usecase

import (
	"errors"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
)

type ProductUseCase struct {
	ProductRepo    sqlite.ProductRepo
	ProductIncRepo sqlite.ProductIncreaseRepo
	ProductDecRepo sqlite.ProductDecreaseRepo
	UserUseCase    UserUseCase
}

func (pus *ProductUseCase) GetProducts() ([]entity.ProductStock, error) {
	result, err := pus.ProductRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pus *ProductUseCase) AddProductValidation(product *entity.Product) error {
	if product.Name == "" {
		return errors.New("name cannot be empty")
	}
	if product.Price == "" {
		return errors.New("price cannot be empty")
	}
	return nil
}

func (pus *ProductUseCase) AddProduct(newProduct *entity.Product) (*entity.Product, error) {
	if err := pus.AddProductValidation(newProduct); err != nil {
		return nil, err
	}
	result, err := pus.ProductRepo.AddProduct(newProduct)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pus *ProductUseCase) EditProductValidation(editedProduct *entity.Product) error {
	if editedProduct.Name == "" {
		return errors.New("name cannot be empty")
	}
	if editedProduct.Price == "" {
		return errors.New("price cannot be empty")
	}
	if editedProduct.ID == 0 {
		return errors.New("ID cannot be 0")
	}
	return nil
}

func (pus *ProductUseCase) EditProduct(editedProduct *entity.Product) (*entity.Product, error) {
	if err := pus.EditProductValidation(editedProduct); err != nil {
		return nil, err
	}
	result, err := pus.ProductRepo.EditProductByID(editedProduct)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (pus *ProductUseCase) AddProductStockValidation(productInc *entity.ProductIncrease) error {
	if productInc.Quantity <= 0 {
		return errors.New("quantity cannot be less than equal 0")
	}
	if productInc.ProductID <= 0 {
		return errors.New("product_id cannot be less than equal 0")
	}
	if productInc.UserID <= 0 {
		return errors.New("user_id cannot be less than equal 0")
	}

	products, err := pus.GetProducts()
	if err != nil {
		return nil
	}
	flagProductID := false
	for _, val := range products {
		if val.ID == productInc.ProductID {
			flagProductID = true
		}
	}
	if !flagProductID {
		return errors.New("product not found")
	}

	users, err := pus.UserUseCase.GetUsers()
	if err != nil {
		return nil
	}
	flagUserID := false
	for _, val := range users {
		if val.ID == productInc.UserID {
			flagUserID = true
		}
	}
	if !flagUserID {
		return errors.New("user not found")
	}

	return nil
}

func (pus *ProductUseCase) AddProductStock(addProductStock *entity.ProductIncrease) (*entity.ProductIncrease, error) {
	if err := pus.AddProductStockValidation(addProductStock); err != nil {
		return nil, err
	}

	result, err := pus.ProductIncRepo.AddProductStock(addProductStock)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pus *ProductUseCase) DecProductValidation(productDec *entity.ProductDec) error {
	var err = errors.New("can't pass the validation")
	if productDec.Quantity < 0 {
		return err
	}
	return nil
}

func (pus *ProductUseCase) DecProductStock(decProductStock *entity.ProductDec) (*entity.ProductDec, error) {
	if err := pus.DecProductValidation(decProductStock); err != nil {
		return nil, err
	}

	result, err := pus.ProductDecRepo.AddDecProductStock(decProductStock)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (pus *ProductUseCase) DeleteProduct(decProductStock int) (string, error) {
	result, err := pus.ProductRepo.DeleteProductByID(decProductStock)
	if err != nil {
		return "", err
	}

	return result, nil
}
