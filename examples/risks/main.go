package main

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
	ctx := context.Background()

	utils.IterateItems(ctx, client.GetRisks)
	reviews := client.RiskReviews()
	utils.IterateItems(ctx, reviews.GetReviews)

	utils.IterateItems(ctx, client.GetThirdPartyRisks)

	reviews = client.ThirdPartyRiskReviews()
	utils.IterateItems(ctx, reviews.GetReviews)

	utils.IterateItems(ctx, client.GetBusinessContinuities)
	reviews = client.BusinessContinuityReviews()
	utils.IterateItems(ctx, reviews.GetReviews)
}
