package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/pkg/model"
)

func (a *Client) GetThirdPartyRisk(ctx context.Context, id int32) (model.ThirdPartyRisk, error) {
	return getDataById[model.ThirdPartyRisk](ctx, "third-party-risks", id, a.getByPath)
}

func (a *Client) GetThirdPartyRisks(ctx context.Context) ([]model.ThirdPartyRisk, error) {
	return getAllData[model.ThirdPartyRisk](ctx, "third-party-risks/index", a.getByPath)
}

func (a *Client) PatchThirdPartyRisk(
	ctx context.Context,
	id int32,
	data *model.ThirdPartyRisk,
) (*model.ThirdPartyRisk, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("third-party-risks/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) GetThirdPartyRiskReviews(ctx context.Context) ([]model.Review, error) {
	return getAllData[model.Review](ctx, "third-party-risk-reviews/index", a.getByPath)
}

func (a *Client) PostThirdPartyRiskReview(ctx context.Context, data *model.Review) (*model.Review, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "third-party-risk-reviews/add", data, a.postOrPatchJsonByPath)
}
