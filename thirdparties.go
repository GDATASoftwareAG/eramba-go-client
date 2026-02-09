package eramba

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetThirdParty(ctx context.Context, id int32) (model.ThirdParty, error) {
	return getDataById[model.ThirdParty](ctx, "third-parties", id, a.getByPath)
}

func (a *Client) GetThirdParties(ctx context.Context) ([]model.ThirdPartyRisk, error) {
	return getAllData[model.ThirdPartyRisk](ctx, "third-parties/index", a.getByPath)
}
