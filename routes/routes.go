package routes

import (
	"github.com/gorilla/mux"
	"github.com/huf0813/pembukuan_tk/ctr"
	"github.com/huf0813/pembukuan_tk/middleware"
	"net/http"
)

type Route struct {
	// user
	HomeCTR    ctr.HomeCTR
	AuthCTR    ctr.AuthCTR
	UserCTR    ctr.UserCTR
	ProductCTR ctr.ProductCTR

	// admin
	AdminCTR ctr.AdminCTR
	Auth     middleware.TokenMiddleware
}

type RouteInterface interface {
	Routes() *mux.Router
}

func (r *Route) Routes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", r.HomeCTR.Welcome).Methods("GET")

	route.HandleFunc("/login", r.AuthCTR.Login).Methods("POST")

	// users
	route.Handle("/dashboard", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.UserCTR.DashboardUser))).Methods("GET")
	route.Handle("/customers/register", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.UserCTR.CustomerRegister))).Methods("POST")
	route.Handle("/customers", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.UserCTR.FetchCustomers))).Methods("GET")
	route.Handle("/products", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.ProductCTR.GetProducts))).Methods("GET")
	route.Handle("/products", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.ProductCTR.AddProduct))).Methods("POST")
	route.Handle("/products/stock", r.Auth.TokenMiddlewareIsUser(http.HandlerFunc(r.ProductCTR.AddProductStock))).Methods("POST")

	// admins
	route.Handle("/admin/dashboard", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.AdminCTR.DashboardAdmin))).Methods("GET")

	return route
}
