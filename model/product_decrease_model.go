package model

type ProductDec struct {
	ID        int `json:"id"`
	ProductID int `json:"product_id"`
	Quantity  int `json:"quantity"`
	InvoiceID int `json:"invoice_id"`
}
