package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetThirdPartyRisk(ctx context.Context, id int32) (model.ThirdPartyRisk, error) {
	return a.getDataById[model.ThirdPartyRisk](ctx, "third-party-risks", id)
}

func (a *Client) GetThirdPartyRisks(ctx context.Context) ([]model.ThirdPartyRisk, error) {
	return a.getAllData[model.ThirdPartyRisk](ctx, "third-party-risks/index")
}

func (a *Client) PostThirdPartyRisk(ctx context.Context, data *model.ThirdPartyRisk) (*model.ThirdPartyRisk, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPost, "third-party-risks/add", data)
}

func (a *Client) PatchThirdPartyRisk(
	ctx context.Context,
	id int32,
	data *model.ThirdPartyRisk,
) (*model.ThirdPartyRisk, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("third-party-risks/%d", id), data)
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
