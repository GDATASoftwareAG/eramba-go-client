package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
)

const (
	RisksTestId = 911
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
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
