package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetSecurityPolicies(ctx context.Context) ([]model.SecurityPolicy, error) {
	return getAllData[model.SecurityPolicy](ctx, "security-policies/index", a.getByPath)
}

func (a *Client) GetSecurityPolicy(ctx context.Context, id int32) (model.SecurityPolicy, error) {
	return getDataById[model.SecurityPolicy](ctx, "security-policies", id, a.getByPath)
}

func (a *Client) PostSecurityPolicy(
	ctx context.Context,
	data *model.SecurityPolicy,
) (*model.SecurityPolicy, error) {
	return postOrPatchJsonByPath(ctx, http.MethodPost, "security-policies/add", data, a.postOrPatchJsonByPath)
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

func (a *Client) GetSecurityPolicyReview(ctx context.Context, id int32) (model.SecurityPolicyReview, error) {
	return getDataById[model.SecurityPolicyReview](ctx, "security-policy-reviews", id, a.getByPath)
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

func (a *Client) DeleteSecurityPolicyReview(ctx context.Context, id int32) error {
	return a.deleteById(ctx, "security-policy-reviews", id)
}
