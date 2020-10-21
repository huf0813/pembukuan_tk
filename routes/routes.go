package routes

import (
	"github.com/gorilla/mux"
	"github.com/huf0813/pembukuan_tk/ctr"
	"github.com/huf0813/pembukuan_tk/middleware"
	"net/http"
)

type Route struct {
	HomeCTR ctr.HomeCTR
	UserCtr ctr.UserCTR
	Auth    middleware.TokenMiddleware
}

type RouteInterface interface {
	Routes() *mux.Router
}

func (r *Route) Routes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", r.HomeCTR.Welcome).Methods("GET")

	route.HandleFunc("/login", r.UserCtr.Login).Methods("POST")
	route.HandleFunc("/users/testing", r.UserCtr.TestingGetUsers).Methods("GET")
	route.Handle("/users", r.Auth.TokenMiddleware(http.HandlerFunc(r.UserCtr.GetUsers))).Methods("GET")

	return route
}
