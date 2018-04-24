package main

import (
	"log"
	"net/http"
	"github.com/n704/go-urlshortener/model"
	"github.com/n704/go-urlshortener/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	model.InitialMigration()
	routes.AddRoutes(myRouter)
	log.Fatal(http.ListenAndServe(":8001", myRouter))
}
