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
	projects, err := client.GetProjects(ctx)
	if err != nil {
		log.Fatalf("Error getting projects: %v", err)
	}
	for i := range projects {
		project := &projects[i]
		fmt.Println(project.Title)

		bytes, err := project.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling project: %v", err)
		}
		fmt.Println(string(bytes))
	}
}
