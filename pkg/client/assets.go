package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/pkg/model"
)

func (a *Client) GetAsset(ctx context.Context, id int32) (model.Asset, error) {
	return getDataById[model.Asset](ctx, "assets", id, a.getByPath)
}

func (a *Client) GetAssets(ctx context.Context) ([]model.Asset, error) {
	return getAllData[model.Asset](ctx, "assets/index", a.getByPath)
}

func (a *Client) PatchAsset(ctx context.Context, id int32, data *model.Asset) (*model.Asset, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("assets/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) GetAssetReviews(ctx context.Context) ([]model.Review, error) {
	return getAllData[model.Review](ctx, "asset-reviews/index", a.getByPath)
}
