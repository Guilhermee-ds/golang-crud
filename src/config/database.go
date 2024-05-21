package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func SetupDatabase() *sql.DB {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name TEXT, email TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
