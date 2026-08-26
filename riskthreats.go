package eramba

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetRiskThreat(ctx context.Context, id int32) (model.RiskThreat, error) {
	return a.getDataById[model.RiskThreat](ctx, "risk-threats", id)
}

func (a *Client) GetRiskThreats(ctx context.Context) ([]model.RiskThreat, error) {
	return a.getAllData[model.RiskThreat](ctx, "risk-threats/index")
}
