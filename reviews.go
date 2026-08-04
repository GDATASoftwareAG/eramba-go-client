package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

type ReviewsClient struct {
	client *Client
	path   string
}

func (a *ReviewsClient) GetReviews(ctx context.Context) ([]model.Review, error) {
	return getAllData[model.Review](ctx, fmt.Sprintf("%s/index", a.path), a.client.getByPath)
}

func (a *ReviewsClient) GetReview(ctx context.Context, id int32) (model.Review, error) {
	return getDataById[model.Review](ctx, a.path, id, a.client.getByPath)
}

func (a *ReviewsClient) DeleteReview(ctx context.Context, id int32) error {
	return a.client.deleteById(ctx, a.path, id)
}

func (a *ReviewsClient) PostReview(ctx context.Context, data *model.Review) (*model.Review, error) {
	return postOrPatchJsonByPath(
		ctx, http.MethodPost,
		fmt.Sprintf("%s/add", a.path),
		data,
		a.client.postOrPatchJsonByPath)
}

func (a *ReviewsClient) PatchReview(ctx context.Context, id int32, data *model.Review) (*model.Review, error) {
	return postOrPatchJsonByPath(
		ctx, http.MethodPatch,
		fmt.Sprintf("%s/%d", a.path, id),
		data,
		a.client.postOrPatchJsonByPath)
}
