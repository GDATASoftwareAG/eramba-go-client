package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetThirdParty(ctx context.Context, id int32) (model.ThirdParty, error) {
	return a.getDataById[model.ThirdParty](ctx, "third-parties", id)
}

func (a *Client) GetThirdParties(ctx context.Context) ([]model.ThirdParty, error) {
	return a.getAllData[model.ThirdParty](ctx, "third-parties/index")
}

func (a *Client) PatchThirdParty(
	ctx context.Context,
	id int32,
	data *model.ThirdParty,
) (*model.ThirdParty, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("third-parties/%d", id), data)
}
