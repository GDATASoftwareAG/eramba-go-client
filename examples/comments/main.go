package main

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/examples/utils"
	"github.com/gdatasoftwareag/eramba-go-client/model"
)

const (
	RisksTestId = 911
)

func main() {
	utils.LoadEnvs()
	client := utils.CreateClientFromEnv()
	ctx := context.Background()
	risks := client.RiskComments()
	utils.IterateItems(ctx, func(ctx context.Context) ([]model.Comment, error) { return risks.GetComments(ctx, RisksTestId) })
}
