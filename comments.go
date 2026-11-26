package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

type CommentsClient struct {
	client *Client
	path   string
	// v2 APIs create comments via POST {path}/{id}/comments instead of {path}/{id}/comments/add
	v2 bool
}

func (a *CommentsClient) GetComments(ctx context.Context, foreignKey int32) ([]model.Comment, error) {
	return a.client.getAllData[model.Comment](ctx, fmt.Sprintf("%s/%d/comments", a.path, foreignKey))
}

func (a *CommentsClient) GetComment(ctx context.Context, foreignKey, id int32) (model.Comment, error) {
	return a.client.getDataById[model.Comment](ctx, fmt.Sprintf("%s/%d/comments", a.path, foreignKey), id)
}

func (a *CommentsClient) DeleteComment(ctx context.Context, foreignKey, id int32) error {
	return a.client.deleteById(ctx, fmt.Sprintf("%s/%d/comments", a.path, foreignKey), id)
}

func (a *CommentsClient) PostComment(ctx context.Context, foreignKey int32, data *model.Comment) (*model.Comment, error) {
	path := fmt.Sprintf("%s/%d/comments", a.path, foreignKey)
	if !a.v2 {
		path += "/add"
	}
	return a.client.postOrPatchJsonByPath(ctx, http.MethodPost, path, data)
}
