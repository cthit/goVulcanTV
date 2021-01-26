package main


import (
	"github.com/joho/godotenv"
	api "github.com/swexbe/govulcantv/internal/api"
	"github.com/swexbe/govulcantv/internal/db"
	"github.com/swexbe/govulcantv/internal/player"
	"log"
)

func main() {
	log.Println("==== Starting goVulcanTV golang backend =====")

	loadDotEnvFile()

	db.Init()
	player.Start()
	api.Init()
}

func loadDotEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Unable to load .env file")
	} else {
		log.Println("Loaded environment variables from .env file")
	}
}
