package entity

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

type ProductStock struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	Stock int64  `json:"stock"`
}

type ProductInsideInvoice struct {
	ProductName       string `json:"product_name"`
	ProductQty        string `json:"product_qty"`
	ProductPrice      string `json:"product_price"`
	ProductTotalPrice string `json:"product_total_price"`
}
