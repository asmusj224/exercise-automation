package database

import (
	"database/sql"
	"log"

	"github.com/asmusj224/exercise-automation/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	database_url := config.Config("DATABASE_URL")

	if database_url == "" {
		log.Fatal("$DATABASE_URL must be set")
	}
	db, err := sql.Open("postgres", database_url)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
