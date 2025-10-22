package model

type ThirdPartyRisk struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	RiskMitigationStrategyId       int32                  `json:"risk_mitigation_strategy_id"`
	RiskAppetiteThresholdAnalysis  *RiskAppetiteThreshold `json:"risk_appetite_threshold_analysis,omitempty"`
	RiskAppetiteThresholdTreatment *RiskAppetiteThreshold `json:"risk_appetite_threshold_treatment,omitempty"`
	Risks1Type2                    RiskClassification     `json:"risk_classifications__third_party_risks_1__type_2"`
	Risks1Type1                    RiskClassification     `json:"risk_classifications__third_party_risks_1__type_1"`
	Risks0Type2                    RiskClassification     `json:"risk_classifications__third_party_risks_0__type_2"`
	Risks0Type1                    RiskClassification     `json:"risk_classifications__third_party_risks_0__type_1"`
	RiskScoreAnalysis              RiskScore              `json:"risk_score_analysis"`
	RiskScoreTreatment             RiskScore              `json:"risk_score_treatment"`
	SecurityServices               SecurityServices       `json:"security_services"`
	Assets                         Assets                 `json:"assets"`
	ThirdParties                   ThirdParties           `json:"third_parties"`
	Threats                        string                 `json:"threats"`
	Tags                           []Tag                  `json:"tags"`
	RiskGrcContacts                []UserOrGroup          `json:"owners"`
	RiskOriginatorContracts        []UserOrGroup          `json:"stakeholders"`
	ThirdPartyReviews              []Review               `json:"third_party_risk_reviews"`

	ThreatTags               []*OnlyId       `json:"threat_tags"`
	VulnerabilityTags        []*OnlyId       `json:"vulnerability_tags"`
	Vulnerabilities          string          `json:"vulnerabilities"`
	RiskExceptions           []*OnlyId       `json:"risk_exceptions"`
	SecurityPoliciesIncident []*OnlyId       `json:"security_policies_incident"`
	Policies                 SecurityPolices `json:"security_policies"`
	PoliciesTreatment        SecurityPolices `json:"security_policies_treatment"`
	Projects                 Projects        `json:"projects"`
	SharedInformation        string          `json:"shared_information"`
	Controlled               string          `json:"controlled"`
}

func (p *ThirdPartyRisk) GetId() int32 {
	return p.Id
}

func (p *ThirdPartyRisk) Link(base string) string {
	return buildLink(base, "third-party-risks", p.Id)
}

var ThirdPartyRiskSkippedFields = []string{
	"id",
	"risk_appetite_threshold_analysis",
	"risk_appetite_threshold_treatment",
	"risk_score_analysis",
	"risk_score_treatment",
	"third_party_risk_reviews",
}

func (p *ThirdPartyRisk) MarshalJSON() ([]byte, error) {
	type Alias ThirdPartyRisk
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, ThirdPartyRiskSkippedFields)
}
