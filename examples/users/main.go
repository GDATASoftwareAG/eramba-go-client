package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
	ctx := context.Background()
	users, err := client.GetUsers(ctx)
	if err != nil {
		log.Fatalf("Error getting users: %v", err)
	}
	for i := range users {
		user := &users[i]
		fmt.Println(user.Id)

		bytes, err := user.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling user: %v", err)
		}
		fmt.Println(string(bytes))
	}
	groups, err := client.GetGroups(ctx)
	if err != nil {
		log.Fatalf("Error getting groups: %v", err)
	}
	for i := range groups {
		group := &groups[i]
		fmt.Println(group.Id)

		bytes, err := group.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling group: %v", err)
		}
		fmt.Println(string(bytes))
	}
}
