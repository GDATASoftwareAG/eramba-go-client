package model

type Risk struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	RiskMitigationStrategyId       int32                  `json:"risk_mitigation_strategy_id"`
	RiskAppetiteThresholdAnalysis  *RiskAppetiteThreshold `json:"risk_appetite_threshold_analysis,omitempty"`
	RiskAppetiteThresholdTreatment *RiskAppetiteThreshold `json:"risk_appetite_threshold_treatment,omitempty"`
	Risks1Type2                    RiskClassification     `json:"risk_classifications__risks_1__type_2"`
	Risks1Type1                    RiskClassification     `json:"risk_classifications__risks_1__type_1"`
	Risks0Type2                    RiskClassification     `json:"risk_classifications__risks_0__type_2"`
	Risks0Type1                    RiskClassification     `json:"risk_classifications__risks_0__type_1"`
	RiskScoreAnalysis              RiskScore              `json:"risk_score_analysis"`
	RiskScoreTreatment             RiskScore              `json:"risk_score_treatment"`
	SecurityServices               SecurityServices       `json:"security_services"`
	Assets                         Assets                 `json:"assets"`
	Threats                        string                 `json:"threats"`
	Tags                           []Tag                  `json:"tags"`
	Reviews                        []Review               `json:"risk_reviews"`
	Policies                       SecurityPolices        `json:"security_policies"`
	PoliciesTreatment              SecurityPolices        `json:"security_policies_treatment"`
	Projects                       Projects               `json:"projects"`
	Owners                         []UserOrGroup          `json:"owners"`
	Stakeholders                   []UserOrGroup          `json:"stakeholders"`
	Review                         string                 `json:"review"`

	ThreatTags               []*OnlyId `json:"threat_tags"`
	VulnerabilityTags        []*OnlyId `json:"vulnerability_tags"`
	Vulnerabilities          string    `json:"vulnerabilities"`
	RiskExceptions           []*OnlyId `json:"risk_exceptions"`
	SecurityPoliciesIncident []*OnlyId `json:"security_policies_incident"`
}

func (p *Risk) GetId() int32 {
	return p.Id
}

func (p *Risk) Link(base string) string {
	return buildLink(base, "risks", p.Id)
}

var RiskSkippedFields = []string{
	"id",
	"risk_appetite_threshold_analysis",
	"risk_appetite_threshold_treatment",
	"risk_score_analysis",
	"risk_score_treatment",
	"risk_reviews",
}

func (p *Risk) MarshalJSON() ([]byte, error) {
	type Alias Risk
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, RiskSkippedFields)
}
