package utils

import (
	"context"
	"encoding/json"
	"fmt"
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

func IterateItems[T any](ctx context.Context, getItems func(ctx context.Context) ([]T, error)) {
	items, err := getItems(ctx)
	if err != nil {
		log.Fatalf("Error getting items: %v", err)
	}
	for i := range items {
		item := &items[i]
		marshaler, ok := any(item).(json.Marshaler)
		if !ok {
			log.Fatalf("%T does not implement json.Marshaler", item)
		}
		bytes, err := marshaler.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling item: %v", err)
		}
		fmt.Println(string(bytes))
	}
}
