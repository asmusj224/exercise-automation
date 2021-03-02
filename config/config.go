package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Config(key string) string {
	err := godotenv.Load()
	if err != nil {
		log.Println(err.Error())
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
