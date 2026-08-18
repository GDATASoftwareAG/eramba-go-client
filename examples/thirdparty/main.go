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
	thirdParties, err := client.GetThirdParties(ctx)
	if err != nil {
		log.Fatalf("Error getting third parties: %v", err)
	}
	for i := range thirdParties {
		group := &thirdParties[i]
		fmt.Println(group.Id)

		bytes, err := group.MarshalJSON()
		if err != nil {
			log.Fatalf("Error marshaling group: %v", err)
		}
		fmt.Println(string(bytes))
	}
}
