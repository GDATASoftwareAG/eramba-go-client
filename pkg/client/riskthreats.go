package client

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/pkg/model"
)

func (a *Client) GetRiskThreat(ctx context.Context, id int32) (model.RiskThreat, error) {
	return getDataById[model.RiskThreat](ctx, "risk-threats", id, a.getByPath)
}

func (a *Client) GetRiskThreats(ctx context.Context) ([]model.RiskThreat, error) {
	return getAllData[model.RiskThreat](ctx, "risk-threats/index", a.getByPath)
}
