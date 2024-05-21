package main

import (
	"github.com/gorilla/mux"
	"github.com/guilhermee-ds/golang-crud/src/config"
	"github.com/guilhermee-ds/golang-crud/src/routes"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	db := config.SetupDatabase()
	defer db.Close()

	router := mux.NewRouter()
	routes.SetupRoutes(router, db)
	log.Fatal(http.ListenAndServe(":8080", config.JsonContentTypeMiddleware(router)))
}
