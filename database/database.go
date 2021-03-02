package database

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/asmusj224/exercise-automation/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() error {
	port, err := ParseDatabasePort()
	if err != nil {
		log.Println("Error parsing database port")
		return err
	}
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))
	if err != nil {
		return err
	}
	DB = db
	return nil
}

func ParseDatabasePort() (uint64, error) {
	var err error
	p := config.Config("DB_PORT")

	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		return 0, err
	}

	return port, nil
}
