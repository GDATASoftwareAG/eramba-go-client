package eramba

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func (a *Client) GetSecurityPolicies(ctx context.Context) ([]model.SecurityPolicy, error) {
	return a.getAllData[model.SecurityPolicy](ctx, "security-policies/index")
}

func (a *Client) GetSecurityPolicy(ctx context.Context, id int32) (model.SecurityPolicy, error) {
	return a.getDataById[model.SecurityPolicy](ctx, "security-policies", id)
}

func (a *Client) PostSecurityPolicy(
	ctx context.Context,
	data *model.SecurityPolicy,
) (*model.SecurityPolicy, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPost, "security-policies/add", data)
}

func (a *Client) PatchSecurityPolicy(
	ctx context.Context,
	id int32,
	data *model.SecurityPolicy,
) (*model.SecurityPolicy, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("security-policies/%d", id), data)
}

func (a *Client) GetSecurityPolicyReviews(ctx context.Context) ([]model.SecurityPolicyReview, error) {
	return a.getAllData[model.SecurityPolicyReview](ctx, "security-policy-reviews/index")
}

func (a *Client) GetSecurityPolicyReview(ctx context.Context, id int32) (model.SecurityPolicyReview, error) {
	return a.getDataById[model.SecurityPolicyReview](ctx, "security-policy-reviews", id)
}

func (a *Client) PatchSecurityPolicyReview(
	ctx context.Context,
	id int32,
	data *model.SecurityPolicyReview,
) (*model.SecurityPolicyReview, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPatch, fmt.Sprintf("security-policy-reviews/%d", id), data)
}

func (a *Client) PostSecurityPolicyReview(
	ctx context.Context,
	data *model.SecurityPolicyReview,
) (*model.SecurityPolicyReview, error) {
	return a.postOrPatchJsonByPath(ctx, http.MethodPost, "security-policy-reviews/add", data)
}

func (a *Client) DeleteSecurityPolicyReview(ctx context.Context, id int32) error {
	return a.deleteById(ctx, "security-policy-reviews", id)
}
