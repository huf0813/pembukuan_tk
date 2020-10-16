package main

import (
	"fmt"
	"net/http"

	"github.com/huf0813/pembukuan_tk/routes"
)

var router routes.Route

func main() {
	const port string = ":8080"
	fmt.Printf("Running at port %s", port)
	if err := http.ListenAndServe(port, router.Routes()); err != nil {
		panic(err)
	}
}
