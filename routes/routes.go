package routes

import (
	"github.com/gorilla/mux"
	"github.com/huf0813/pembukuan_tk/ctr"
	"github.com/huf0813/pembukuan_tk/middleware"
	"net/http"
)

type Route struct {
	HomeCTR     ctr.HomeCTR
	AuthCTR     ctr.AuthCTR
	UserCTR     ctr.UserCTR
	ProductCTR  ctr.ProductCTR
	InvoiceCTR  ctr.InvoiceCTR
	CustomerCTR ctr.CustomerCTR
	AdminCTR    ctr.AdminCTR
	Auth        middleware.TokenMiddleware
}

type RouteInterface interface {
	Routes() *mux.Router
}

func (r *Route) Routes() *mux.Router {
	route := mux.NewRouter()

	// welcome to API
	route.HandleFunc("/", r.HomeCTR.Welcome).Methods("GET")

	// auth
	route.HandleFunc("/login", r.AuthCTR.Login).Methods("POST")
	route.HandleFunc("/token/expired", r.AuthCTR.CheckToken).Methods("POST")

	// users
	route.Handle("/dashboard", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.UserCTR.DashboardUser))).Methods("GET")
	route.Handle("/customers", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.CustomerCTR.CustomerRegister))).Methods("POST")
	route.Handle("/customers", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.CustomerCTR.FetchCustomers))).Methods("GET")
	route.Handle("/customers", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.CustomerCTR.EditCustomer))).Methods("PUT")
	route.Handle("/products", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.ProductCTR.GetProducts))).Methods("GET")
	route.Handle("/products", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.ProductCTR.AddProduct))).Methods("POST")
	route.Handle("/products", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.ProductCTR.EditProduct))).Methods("PUT")
	route.Handle("/products/stock", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.ProductCTR.AddProductStock))).Methods("POST")
	route.Handle("/invoice", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.InvoiceCTR.AddNewInvoice))).Methods("POST")
	route.Handle("/invoice", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.InvoiceCTR.GetInvoices))).Methods("GET")
	route.Handle("/invoice/detail", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.InvoiceCTR.GetInvoiceByID))).Methods("POST")
	route.Handle("/statistics", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.InvoiceCTR.GetStatistics))).Methods("POST")

	// admins
	route.Handle("/admin/dashboard", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.AdminCTR.DashboardAdmin))).Methods("GET")
	route.Handle("/admin/customers", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.CustomerCTR.CustomerRegister))).Methods("POST")
	route.Handle("/admin/customers", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.CustomerCTR.FetchCustomers))).Methods("GET")
	route.Handle("/admin/customers", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.CustomerCTR.EditCustomer))).Methods("PUT")
	route.Handle("/admin/products", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.ProductCTR.GetProducts))).Methods("GET")
	route.Handle("/admin/products", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.ProductCTR.AddProduct))).Methods("POST")
	route.Handle("/admin/products", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.ProductCTR.EditProduct))).Methods("PUT")
	route.Handle("/admin/products/stock", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.ProductCTR.AddProductStock))).Methods("POST")
	route.Handle("/admin/invoice", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.InvoiceCTR.AddNewInvoice))).Methods("POST")
	route.Handle("/admin/invoice", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.InvoiceCTR.GetInvoices))).Methods("GET")
	route.Handle("/admin/invoice/detail", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.InvoiceCTR.GetInvoiceByID))).Methods("POST")
	route.Handle("/admin/users", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.UserCTR.FetchUsers))).Methods("GET")
	route.Handle("/admin/users", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.UserCTR.AddUser))).Methods("POST")
	route.Handle("/admin/users", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.UserCTR.EditedUser))).Methods("PUT")
	route.Handle("/admin/statistics", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.InvoiceCTR.GetStatistics))).Methods("POST")

	return route
}
