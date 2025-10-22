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
	users, err := client.GetUsers(ctx)
	if err != nil {
		log.Fatalf("Error getting users: %v", err)
	}
	for i := range users {
		user := &users[i]
		fmt.Println(user.Id)

		bla, err := user.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling user: %v", err)
		}
		fmt.Println(string(bla))
	}
	groups, err := client.GetGroups(ctx)
	if err != nil {
		log.Fatalf("Error getting groups: %v", err)
	}
	for i := range groups {
		group := &groups[i]
		fmt.Println(group.Id)

		bla, err := group.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling group: %v", err)
		}
		fmt.Println(string(bla))
	}
}
