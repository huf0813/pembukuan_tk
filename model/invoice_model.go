package model

type Invoice struct {
	ID         int `json:"id"`
	CustomerID int `json:"customer_id"`
	UserID     int `json:"user_id"`
}

type ProductDecReq struct {
	ProductID int `json:"product_id"`
	Qty       int `json:"qty"`
}

type InvoiceReq struct {
	CustomerID  int             `json:"customer_id"`
	UserID      int             `json:"user_id"`
	ListProduct []ProductDecReq `json:"list_product"`
}
