package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/huf0813/pembukuan_tk/routes"
)

var router routes.Route

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Printf("Running at port %s", port)
	if err := http.ListenAndServe(port, router.Routes()); err != nil {
		panic(err)
	}
}
