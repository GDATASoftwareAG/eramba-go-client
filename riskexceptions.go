package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetRiskException(ctx context.Context, id int32) (model.RiskException, error) {
	return a.getDataById[model.RiskException](ctx, "risk-exceptions", id)
}

func (a *Client) GetRiskExceptions(ctx context.Context) ([]model.RiskException, error) {
	return a.getAllData[model.RiskException](ctx, "risk-exceptions/index")
}

func (a *Client) PostRiskException(ctx context.Context, data *model.RiskException) (*model.RiskException, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPost, "risk-exceptions/add", data)
}

func (a *Client) PatchRiskException(
	ctx context.Context,
	id int32,
	data *model.RiskException,
) (*model.RiskException, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("risk-exceptions/%d", id), data)
}
