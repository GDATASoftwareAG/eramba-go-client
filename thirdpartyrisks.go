package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetThirdPartyRisk(ctx context.Context, id int32) (model.ThirdPartyRisk, error) {
	return getDataById[model.ThirdPartyRisk](ctx, "third-party-risks", id, a.getByPath)
}

func (a *Client) GetThirdPartyRisks(ctx context.Context) ([]model.ThirdPartyRisk, error) {
	return getAllData[model.ThirdPartyRisk](ctx, "third-party-risks/index", a.getByPath)
}

func (a *Client) PostThirdPartyRisk(ctx context.Context, data *model.ThirdPartyRisk) (*model.ThirdPartyRisk, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "third-party-risks/add", data, a.postOrPatchJsonByPath)
}

func (a *Client) PatchThirdPartyRisk(
	ctx context.Context,
	id int32,
	data *model.ThirdPartyRisk,
) (*model.ThirdPartyRisk, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("third-party-risks/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) ThirdPartyRiskComments() *CommentsClient {
	return &CommentsClient{
		client: a,
		path:   "third-party-risks",
	}
}

func (a *Client) ThirdPartyRiskReviews() *ReviewsClient {
	return &ReviewsClient{
		client: a,
		path:   "third-party-risk-reviews",
	}
}
