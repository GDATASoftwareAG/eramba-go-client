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
}

func (a *CommentsClient) GetComments(ctx context.Context, foreignKey int32) ([]model.Comment, error) {
	return getAllData[model.Comment](ctx, fmt.Sprintf("%s/%d/comments", a.path, foreignKey), a.client.getByPath)
}

func (a *CommentsClient) GetComment(ctx context.Context, foreignKey, id int32) (model.Comment, error) {
	return getDataById[model.Comment](ctx, fmt.Sprintf("%s/%d/comments", a.path, foreignKey), id, a.client.getByPath)
}

func (a *CommentsClient) DeleteComment(ctx context.Context, foreignKey, id int32) error {
	return a.client.deleteById(ctx, fmt.Sprintf("%s/%d/comments", a.path, foreignKey), id)
}

func (a *CommentsClient) PostComment(ctx context.Context, foreignKey int32, data *model.Comment) (*model.Comment, error) {
	return postOrPatchJsonByPath(
		ctx, http.MethodPost,
		fmt.Sprintf("%s/%d/comments/add", a.path, foreignKey),
		data,
		a.client.postOrPatchJsonByPath)
}
