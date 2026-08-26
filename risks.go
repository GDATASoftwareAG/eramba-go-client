package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetRisk(ctx context.Context, id int32) (model.Risk, error) {
	return a.getDataById[model.Risk](ctx, "risks", id)
}

func (a *Client) GetRisks(ctx context.Context) ([]model.Risk, error) {
	return a.getAllData[model.Risk](ctx, "risks/index")
}

func (a *Client) PostRisk(ctx context.Context, data *model.Risk) (*model.Risk, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPost, "risks/add", data)
}

func (a *Client) PatchRisk(ctx context.Context, id int32, data *model.Risk) (*model.Risk, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("risks/%d", id), data)
}

func (a *Client) RiskComments() *CommentsClient {
	return &CommentsClient{
		client: a,
		path:   "risks",
	}
}

func (a *Client) RiskReviews() *ReviewsClient {
	return &ReviewsClient{
		client: a,
		path:   "risk-reviews",
	}
}
