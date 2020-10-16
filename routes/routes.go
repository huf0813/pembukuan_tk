package routes

import (
	"github.com/gorilla/mux"
	"github.com/huf0813/pembukuan_tk/ctr"
)

type Route struct {
	HomeCTR ctr.HomeCTR
}

type RouteInterface interface {
	Routes() *mux.Router
}

func (r *Route) Routes() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", r.HomeCTR.Welcome).Methods("GET")

	return route
}
