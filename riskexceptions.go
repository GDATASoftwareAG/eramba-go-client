package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetRiskException(ctx context.Context, id int32) (model.RiskException, error) {
	return getDataById[model.RiskException](ctx, "risk-exceptions", id, a.getByPath)
}

func (a *Client) GetRiskExceptions(ctx context.Context) ([]model.RiskException, error) {
	return getAllData[model.RiskException](ctx, "risk-exceptions/index", a.getByPath)
}

func (a *Client) PostRiskException(ctx context.Context, data *model.RiskException) (*model.RiskException, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "risk-exceptions/add", data, a.postOrPatchJsonByPath)
}

func (a *Client) PatchRiskException(
	ctx context.Context,
	id int32,
	data *model.RiskException,
) (*model.RiskException, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("risk-exceptions/%d", id), data, a.postOrPatchJsonByPath)
}
