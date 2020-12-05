package entity

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

type InvoiceWithDetail struct {
	ID                int                    `json:"id"`
	CustomerName      string                 `json:"customer_name"`
	CustomerPhone     string                 `json:"customer_phone"`
	CustomerEmail     string                 `json:"customer_email"`
	CustomerAddress   string                 `json:"customer_address"`
	TotalInvoicePrice string                 `json:"total_invoice_price"`
	CreatedAt         string                 `json:"created_at"`
	UpdatedAt         string                 `json:"updated_at"`
	Products          []ProductInsideInvoice `json:"products"`
}
