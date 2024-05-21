package routes

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/guilhermee-ds/golang-crud/src/controller"
)

func SetupRoutes(router *mux.Router, db *sql.DB) {
	router.HandleFunc("/users", controllers.GetUsers(db)).Methods("GET")
	router.HandleFunc("/users/{id}", controllers.GetUser(db)).Methods("GET")
	router.HandleFunc("/users", controllers.CreateUser(db)).Methods("POST")
	router.HandleFunc("/users/{id}", controllers.UpdateUser(db)).Methods("PUT")
	router.HandleFunc("/users/{id}", controllers.DeleteUser(db)).Methods("DELETE")
}
