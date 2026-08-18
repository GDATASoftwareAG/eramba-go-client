package main

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
	ctx := context.Background()
	utils.IterateItems(ctx, client.GetUsers)
	utils.IterateItems(ctx, client.GetGroups)
}
