package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]

		var u User
		err := db.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Email)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
	}
}
