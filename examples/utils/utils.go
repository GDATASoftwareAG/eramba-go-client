package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	erambaClient "github.com/gdatasoftwareag/eramba-go-client"
)

func LoadEnvs() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
}

func CreateClientFromEnv() erambaClient.Client {
	erambaPassword := os.Getenv("ERAMBA_PASSWORD")
	erambaUser := os.Getenv("ERAMBA_USER")
	erambaUrl := os.Getenv("ERAMBA_URL")
	client := erambaClient.New(erambaUrl, erambaUser, erambaPassword)
	return client
}
