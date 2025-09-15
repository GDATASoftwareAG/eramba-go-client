package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetRisk(ctx context.Context, id int32) (model.Risk, error) {
	return getDataById[model.Risk](ctx, "risks", id, a.getByPath)
}

func (a *Client) GetRisks(ctx context.Context) ([]model.Risk, error) {
	return getAllData[model.Risk](ctx, "risks/index", a.getByPath)
}

func (a *Client) PostRisk(ctx context.Context, data *model.Risk) (*model.Risk, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "risks/add", data, a.postOrPatchJsonByPath)
}

func (a *Client) PatchRisk(ctx context.Context, id int32, data *model.Risk) (*model.Risk, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("risks/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) GetRiskReviews(ctx context.Context) ([]model.Review, error) {
	return getAllData[model.Review](ctx, "risk-reviews/index", a.getByPath)
}

func (a *Client) PostRiskReview(ctx context.Context, data *model.Review) (*model.Review, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "risk-reviews/add", data, a.postOrPatchJsonByPath)
}

func (a *Client) PatchRiskReview(ctx context.Context, id int32, data *model.Review) (*model.Review, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("risk-reviews/%d", id), data, a.postOrPatchJsonByPath)
}
