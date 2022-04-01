package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	DATABASE_URL string
	TWITTER_KEY string
	TWITTER_SECRET string
	DATABASE_NAME string
	KAKAO_CLIENT_ID string
)

func GetENV() {
	err := godotenv.Load(".env")
	
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	
	DATABASE_NAME = os.Getenv("DATABASE_NAME")
	DATABASE_URL = os.Getenv("DATABASE_URL")
	TWITTER_KEY = os.Getenv("TWITTER_KEY")
	TWITTER_SECRET = os.Getenv("TWITTER_SECRET")
	KAKAO_CLIENT_ID = os.Getenv("KAKAO_CLIENT_ID")

	fmt.Print(DATABASE_NAME, DATABASE_URL)
}