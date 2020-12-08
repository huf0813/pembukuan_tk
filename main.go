package main

import (
	"fmt"
	ConnInit "github.com/huf0813/pembukuan_tk/db/sqlite"
	"github.com/huf0813/pembukuan_tk/entity"
	"github.com/huf0813/pembukuan_tk/repository/sqlite"
	"github.com/huf0813/pembukuan_tk/routes"
	"github.com/huf0813/pembukuan_tk/utils"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

var (
	router         routes.Route
	sqliteConnInit ConnInit.ConnSqlite
	hash           utils.Hashing

	userRepo     sqlite.UserRepo
	userTypeRepo sqlite.UserTypeRepo
)

func init() {
	if err := sqliteConnInit.AutoDropDB(); err != nil {
		panic(err)
	}

	admin := "admin"
	if err := userTypeRepo.AddUserType(admin); err != nil {
		panic(err)
	}
	user := "user"
	if err := userTypeRepo.AddUserType(user); err != nil {
		panic(err)
	}

	jo := "jo"
	joPass, err := hash.HashPass(jo)
	if err != nil {
		panic(err)
	} else {
		if _, err := userRepo.AddUser(&entity.User{
			Username:   jo,
			Password:   string(joPass),
			UserTypeID: 2,
		}); err != nil {
			panic(err)
		}
	}
	har := "har"
	harPass, err := hash.HashPass(har)
	if err != nil {
		panic(err)
	} else {
		if _, err := userRepo.AddUser(&entity.User{
			Username:   har,
			Password:   string(harPass),
			UserTypeID: 1,
		}); err != nil {
			panic(err)
		}
	}
}

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
