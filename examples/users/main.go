package main

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
	ctx := context.Background()
	utils.IterateItems(ctx, client.Users().GetAll)
	utils.IterateItems(ctx, client.Groups().GetAll)
}
