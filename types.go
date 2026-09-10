package eramba

import "github.com/gdatasoftwareag/eramba-go-client/model"

func (a *Client) Assets() *GetAndPatchClientWithCommentAndReview[model.Asset] {
	return &GetAndPatchClientWithCommentAndReview[model.Asset]{
		client:     a,
		path:       "assets",
		pathReview: "asset-reviews",
		model:      "Assets",
	}
}

func (a *Client) BusinessContinuities() *GetAndPatchClientWithCommentAndReview[model.BusinessContinuity] {
	return &GetAndPatchClientWithCommentAndReview[model.BusinessContinuity]{
		client:     a,
		path:       "business-continuities",
		pathReview: "business-continuity-reviews",
		model:      "BusinessContinuities",
	}
}

func (a *Client) Groups() *GetClient[model.Group] {
	return &GetClient[model.Group]{
		client: a,
		path:   "groups",
	}
}

func (a *Client) Risks() *GetAndPatchClientWithCommentAndReview[model.Risk] {
	return &GetAndPatchClientWithCommentAndReview[model.Risk]{
		client:     a,
		path:       "risks",
		pathReview: "risk-reviews",
		model:      "Risks",
	}
}

func (a *Client) ThirdPartyRisks() *GetAndPatchClientWithCommentAndReview[model.ThirdPartyRisk] {
	return &GetAndPatchClientWithCommentAndReview[model.ThirdPartyRisk]{
		client:     a,
		path:       "third-party-risks",
		pathReview: "third-party-risk-reviews",
		model:      "ThirdPartyRisks",
	}
}

func (a *Client) Users() *GetClient[model.User] {
	return &GetClient[model.User]{
		client: a,
		path:   "users",
	}
}

func (a *Client) ThirdParties() *GetAndPatchClient[model.ThirdParty] {
	return &GetAndPatchClient[model.ThirdParty]{
		client: a,
		path:   "third-parties",
	}
}

func (a *Client) SecurityServices() *GetAndPatchClientWithComment[model.SecurityService] {
	return &GetAndPatchClientWithComment[model.SecurityService]{
		client: a,
		path:   "security-services",
	}
}

func (a *Client) RiskThreats() *GetClient[model.RiskThreat] {
	return &GetClient[model.RiskThreat]{
		client: a,
		path:   "risk-threats",
	}
}

func (a *Client) RiskExceptions() *GetAndPatchClient[model.RiskException] {
	return &GetAndPatchClient[model.RiskException]{
		client: a,
		path:   "risk-exceptions",
	}
}

func (a *Client) SecurityPolicies() *GetAndPatchClient[model.SecurityPolicy] {
	return &GetAndPatchClient[model.SecurityPolicy]{
		client: a,
		path:   "security-policies",
	}
}

func (a *Client) SecurityPolicyReviews() *GetAndPatchClient[model.SecurityPolicyReview] {
	return &GetAndPatchClient[model.SecurityPolicyReview]{
		client: a,
		path:   "security-policy-reviews",
	}
}

func (a *Client) Projects() *GetAndPatchClientWithComment[model.Project] {
	return &GetAndPatchClientWithComment[model.Project]{
		client: a,
		path:   "projects",
	}
}
