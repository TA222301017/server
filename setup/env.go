package setup

import (
	"log"

	"github.com/joho/godotenv"
)

func Env() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Failed to load environment variables : %s\n", err)
	}
	log.Println("Loaded environment variables")
}
