package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/pkg/model"
)

func (a *Client) GetSecurityPolicies(ctx context.Context) ([]model.SecurityPolicy, error) {
	return getAllData[model.SecurityPolicy](ctx, "security-policies/index", a.getByPath)
}

func (a *Client) PatchSecurityPolicy(
	ctx context.Context,
	id int32,
	data *model.SecurityPolicy,
) (*model.SecurityPolicy, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("security-policies/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) GetSecurityPolicyReviews(ctx context.Context) ([]model.SecurityPolicyReview, error) {
	return getAllData[model.SecurityPolicyReview](ctx, "security-policy-reviews/index", a.getByPath)
}

func (a *Client) PatchSecurityPolicyReview(
	ctx context.Context,
	id int32,
	data *model.SecurityPolicyReview,
) (*model.SecurityPolicyReview, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("security-policy-reviews/%d", id), data, a.postOrPatchJsonByPath)
}

func (a *Client) PostSecurityPolicyReview(
	ctx context.Context,
	data *model.SecurityPolicyReview,
) (*model.SecurityPolicyReview, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "security-policy-reviews/add", data, a.postOrPatchJsonByPath)
}
