package main

import (
	"fmt"
	"github.com/huf0813/pembukuan_tk/routes"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var router routes.Route

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	fmt.Printf("Running at port %s", ":"+port)
	if err := http.ListenAndServe(":"+port, router.Routes()); err != nil {
		panic(err)
	}
}
