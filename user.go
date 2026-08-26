package eramba

import (
	"context"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetUser(ctx context.Context, id int32) (model.User, error) {
	return a.getDataById[model.User](ctx, "users", id)
}

func (a *Client) GetUsers(ctx context.Context) ([]model.User, error) {
	return a.getAllData[model.User](ctx, "users/index")
}
