package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetSecurityService(ctx context.Context, id int32) (model.SecurityService, error) {
	return a.getDataById[model.SecurityService](ctx, "security-services", id)
}

func (a *Client) GetSecurityServices(ctx context.Context) ([]model.SecurityService, error) {
	return a.getAllData[model.SecurityService](ctx, "security-services/index")
}

func (a *Client) PatchSecurityService(
	ctx context.Context,
	id int32,
	data *model.SecurityService,
) (*model.SecurityService, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("security-services/%d", id), data)
}

func (a *Client) SecurityServiceComments() *CommentsClient {
	return &CommentsClient{
		client: a,
		path:   "security-services",
	}
}
