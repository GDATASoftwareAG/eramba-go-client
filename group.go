package eramba

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetGroup(ctx context.Context, id int32) (model.Group, error) {
	return getDataById[model.Group](ctx, "groups", id, a.getByPath)
}

func (a *Client) GetGroups(ctx context.Context) ([]model.Group, error) {
	return getAllData[model.Group](ctx, "groups/index", a.getByPath)
}
