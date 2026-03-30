package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	erambaClient "github.com/gdatasoftwareag/eramba-go-client"
)

const (
	RisksTestId = 911
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
	risks := client.RiskComments()

	comments, err := risks.GetComments(ctx, RisksTestId)
	if err != nil {
		log.Fatal(err)
	}
	for i := range comments {
		comment := &comments[i]
		fmt.Println(comment.Message)

		bytes, err := comment.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling project: %v", err)
		}
		fmt.Println(string(bytes))
	}
}
