package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetAsset(ctx context.Context, id int32) (model.Asset, error) {
	return a.getDataById[model.Asset](ctx, "assets", id)
}

func (a *Client) GetAssets(ctx context.Context) ([]model.Asset, error) {
	return a.getAllData[model.Asset](ctx, "assets/index")
}

func (a *Client) PatchAsset(ctx context.Context, id int32, data *model.Asset) (*model.Asset, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("assets/%d", id), data)
}

func (a *Client) PostAsset(ctx context.Context, data *model.Asset) (*model.Asset, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPost, "assets/add", data)
}

func (a *Client) DeleteAsset(ctx context.Context, id int32) error {
	return a.deleteById(ctx, "assets", id)
}

func (a *Client) AssetComments() *CommentsClient {
	return &CommentsClient{
		client: a,
		path:   "assets",
	}
}

func (a *Client) AssetReviews() *ReviewsClient {
	return &ReviewsClient{
		client: a,
		path:   "asset-reviews",
	}
}
