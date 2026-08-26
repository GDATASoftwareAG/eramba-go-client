package eramba

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetGroup(ctx context.Context, id int32) (model.Group, error) {
	return a.getDataById[model.Group](ctx, "groups", id)
}

func (a *Client) GetGroups(ctx context.Context) ([]model.Group, error) {
	return a.getAllData[model.Group](ctx, "groups/index")
}
