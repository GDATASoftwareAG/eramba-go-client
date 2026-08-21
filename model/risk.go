package model

import (
	"encoding/json"
)

type Risk struct {
	Id          int32  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`

	RiskMitigationStrategyId       int32                        `json:"risk_mitigation_strategy_id"`
	RiskAppetiteThresholdAnalysis  *RiskAppetiteThreshold       `json:"risk_appetite_threshold_analysis,omitempty"`
	RiskAppetiteThresholdTreatment *RiskAppetiteThreshold       `json:"risk_appetite_threshold_treatment,omitempty"`
	RiskClassificationAnalysis     map[int32]RiskClassification `json:"-"`
	RiskClassificationTreatment    map[int32]RiskClassification `json:"-"`
	RiskScoreAnalysis              RiskScore                    `json:"risk_score_analysis"`
	RiskScoreTreatment             RiskScore                    `json:"risk_score_treatment"`
	SecurityServices               SecurityServices             `json:"security_services"`
	Assets                         Assets                       `json:"assets"`
	Threats                        string                       `json:"threats"`
	Tags                           []Tag                        `json:"tags"`
	RiskGrcContacts                []UserOrGroup                `json:"owners"`
	RiskOriginatorContacts         []UserOrGroup                `json:"stakeholders"`
	Reviews                        []Review                     `json:"risk_reviews"`

	ThreatTags               []*OnlyId       `json:"threat_tags"`
	VulnerabilityTags        []*OnlyId       `json:"vulnerability_tags"`
	Vulnerabilities          string          `json:"vulnerabilities"`
	RiskExceptions           []*OnlyId       `json:"risk_exceptions"`
	SecurityPoliciesIncident []*OnlyId       `json:"security_policies_incident"`
	Policies                 SecurityPolices `json:"security_policies"`
	PoliciesTreatment        SecurityPolices `json:"security_policies_treatment"`
	Projects                 Projects        `json:"projects"`

	CustomFields CustomFields `json:"-"`
}

func (p *Risk) GetId() int32 {
	return p.Id
}

func (p *Risk) Link(base string) string {
	return ErambaViewLink(base, "risks", p.Id)
}

var RiskSkippedFields = []string{
	FieldId,
	FieldRiskAppetiteThresholdAnalysis,
	FieldRiskAppetiteThresholdTreatment,
	FieldRiskScoreAnalysis,
	FieldRiskScoreTreatment,
	"risk_reviews",
}

func (p *Risk) UnmarshalJSON(data []byte) error {
	type Alias Risk // avoid recursion
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &aux.Alias); err != nil {
		return err
	}
	customFields, err := UnmarshalCustomFields(data)
	if err != nil {
		return err
	}
	p.CustomFields = customFields

	rcAnalysis, rcTreatment, err := UnmarshalRiskClassification(FieldRiskClassificationsRisksPrefix, data)
	if err != nil {
		return err
	}
	p.RiskClassificationAnalysis = rcAnalysis
	p.RiskClassificationTreatment = rcTreatment

	return nil
}

func (p *Risk) MarshalJSON() ([]byte, error) {
	type Alias Risk
	aux := Alias(*p)
	extraFields := MarshalRiskClassification(
		FieldRiskClassificationsRisksPrefix, p.RiskClassificationAnalysis, p.RiskClassificationTreatment)
	return MarshalWithSpecialFields(aux, p.CustomFields, extraFields, RiskSkippedFields)
}
