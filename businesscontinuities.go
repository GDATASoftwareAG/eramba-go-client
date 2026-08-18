package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetBusinessContinuity(ctx context.Context, id int32) (model.BusinessContinuity, error) {
	return getDataById[model.BusinessContinuity](ctx, "business-continuities", id, a.getByPath)
}

func (a *Client) GetBusinessContinuities(ctx context.Context) ([]model.BusinessContinuity, error) {
	return getAllData[model.BusinessContinuity](ctx, "business-continuities/index", a.getByPath)
}

func (a *Client) PostBusinessContinuity(ctx context.Context, data *model.BusinessContinuity) (*model.BusinessContinuity, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "business-continuities/add", data, a.postOrPatchJsonByPath)
}

func (a *Client) PatchBusinessContinuity(
	ctx context.Context,
	id int32,
	data *model.BusinessContinuity,
) (*model.BusinessContinuity, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("business-continuities/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) BusinessContinuityComments() *CommentsClient {
	return &CommentsClient{
		client: a,
		path:   "business-continuities",
	}
}

func (a *Client) BusinessContinuityReviews() *ReviewsClient {
	return &ReviewsClient{
		client: a,
		path:   "business-continuity-reviews",
	}
}
