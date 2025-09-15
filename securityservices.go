package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetSecurityService(ctx context.Context, id int32) (model.SecurityService, error) {
	return getDataById[model.SecurityService](ctx, "security-services", id, a.getByPath)
}

func (a *Client) GetSecurityServices(ctx context.Context) ([]model.SecurityService, error) {
	return getAllData[model.SecurityService](ctx, "security-services/index", a.getByPath)
}

func (a *Client) PatchSecurityService(
	ctx context.Context,
	id int32,
	data *model.SecurityService,
) (*model.SecurityService, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("security-services/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) GetSecurityServiceComments(ctx context.Context, id int32) ([]model.Comment, error) {
	return getAllData[model.Comment](ctx, fmt.Sprintf("security-services/%d/comments", id), a.getByPath)
}
