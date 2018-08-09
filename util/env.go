package util

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func GetEnv(key string) string {
	value := os.Getenv(key)
	return value
}

func LoadEnv() {
	filename := ".env"
	err := godotenv.Load(filename)
	if err != nil {
		log.Fatal(".env file not loaded")
	}
}
