package model

type Product struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Price         string `json:"price"`
	ProductTypeID int    `json:"product_type_id"`
}
