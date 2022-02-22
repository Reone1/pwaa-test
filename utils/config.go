package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL string
	TWITTER_KEY string
	TWITTER_SECRET string
	DATABASE_NAME string
)

func GetENV() {
	err := godotenv.Load()
	
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	DATABASE_NAME = os.Getenv("DATABASE_NAME")
	DATABASE_URL = os.Getenv("DATABASE_URL")
	TWITTER_KEY = os.Getenv("TWITTER_KEY")
	TWITTER_SECRET = os.Getenv("TWITTER_SECRET")
}