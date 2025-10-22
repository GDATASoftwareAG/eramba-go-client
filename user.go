package eramba

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetUser(ctx context.Context, id int32) (model.User, error) {
	return getDataById[model.User](ctx, "users", id, a.getByPath)
}

func (a *Client) GetUsers(ctx context.Context) ([]model.User, error) {
	return getAllData[model.User](ctx, "users/index", a.getByPath)
}
