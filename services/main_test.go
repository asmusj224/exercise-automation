package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "admin", "password", "db_test"))
	if err != nil {
		log.Fatal("Unable to connect to db", err)
	}

	testQueries = New(db)

	os.Exit(m.Run())
}
