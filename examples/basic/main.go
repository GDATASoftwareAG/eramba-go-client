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
	risks, err := client.GetBusinessContinuities(ctx)
	if err != nil {
		log.Fatalf("Error getting risks: %v", err)
	}
	for i := range risks {
		risk := &risks[i]
		fmt.Println(risk.Title)

		bytes, err := risk.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling project: %v", err)
		}
		fmt.Println(string(bytes))
	}
}
