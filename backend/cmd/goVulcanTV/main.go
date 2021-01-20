package main


import (
	"github.com/joho/godotenv"
	"github.com/swexbe/govulcantv/internal/db"
	"log"
)

func main() {
	log.Println("==== Starting goVulcanTV golang backend =====")

	loadDotEnvFile()

	db.Init()
}

func loadDotEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file")
	} else {
		log.Println("Loaded environment variables from .env file")
	}
}
