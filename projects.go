package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetProjects(ctx context.Context) ([]model.Project, error) {
	return getAllData[model.Project](ctx, "projects/index", a.getByPath)
}

func (a *Client) PostProject(ctx context.Context, data *model.Project) (*model.Project, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "projects/add", data, a.postOrPatchJsonByPath)
}

func (a *Client) PatchProject(ctx context.Context, id int32, data *model.Project) (*model.Project, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("projects/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) PostProjectComment(ctx context.Context, id int32, data *model.Comment) (*model.Comment, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, fmt.Sprintf("projects/%d/comments/add", id), data, a.postOrPatchJsonByPath)
}
