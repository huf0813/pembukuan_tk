package routes

import (
	"github.com/gorilla/mux"
	"github.com/huf0813/pembukuan_tk/ctr"
	"github.com/huf0813/pembukuan_tk/middleware"
	"net/http"
)

type Route struct {
	HomeCTR  ctr.HomeCTR
	AuthCTR  ctr.AuthCTR
	UserCTR  ctr.UserCTR
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

	// admins
	route.Handle("/admin/dashboard", r.Auth.TokenMiddlewareIsAdmin(http.HandlerFunc(r.AdminCTR.DashboardAdmin))).Methods("GET")
	return route
}
