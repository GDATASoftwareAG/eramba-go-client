package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	erambaClient "github.com/gdatasoftwareag/eramba-go-client"
)

func main() {
	if _, err := os.Stat(".env"); err == nil {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	erambaPassword := os.Getenv("ERAMBA_PASSWORD")
	erambaUser := os.Getenv("ERAMBA_USER")
	erambaUrl := os.Getenv("ERAMBA_URL")
	client := erambaClient.New(erambaUrl, erambaUser, erambaPassword)
	ctx := context.Background()
	risks, err := client.GetRisks(ctx)
	if err != nil {
		log.Fatalf("Error getting risks: %v", err)
	}
	for i := range risks {
		risk := &risks[i]
		fmt.Println(risk.Title)

		bla, err := risk.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling project: %v", err)
		}
		fmt.Println(string(bla))
	}
}
