package main

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
	ctx := context.Background()

	utils.IterateItems(ctx, client.Risks().GetAll)
	reviews := client.Risks().Reviews()
	utils.IterateItems(ctx, reviews.GetAll)

	utils.IterateItems(ctx, client.ThirdPartyRisks().GetAll)

	reviews = client.ThirdPartyRisks().Reviews()
	utils.IterateItems(ctx, reviews.GetAll)

	utils.IterateItems(ctx, client.BusinessContinuities().GetAll)
	reviews = client.BusinessContinuities().Reviews()
	utils.IterateItems(ctx, reviews.GetAll)
}
