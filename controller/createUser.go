package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u User

		json.NewDecoder(r.Body).Decode(&u)

		err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", u.Name, u.Email).Scan(&u.ID)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(u)
	}
}
