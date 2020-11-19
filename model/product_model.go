package model

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	ProductTypeID int    `json:"product_type_id"`
}

type ProductStockAndType struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Price       string `json:"price"`
	ProductType string `json:"product_type"`
	Stock       int64  `json:"stock"`
}
