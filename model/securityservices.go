package model

import "encoding/json"

type SecurityServiceType int32

const (
	SecurityServiceTypeDesign     SecurityServiceType = 2
	SecurityServiceTypeProduction SecurityServiceType = 4
)

type SecurityServices []*SecurityService

func (p SecurityServices) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type SecurityService struct {
	Id                      int32               `json:"id"`
	Title                   string              `json:"name"`
	Description             string              `json:"objective"`
	SecurityServiceType     SecurityServiceType `json:"security_service_type_id"`
	ControlOperatorContacts []UserOrGroup       `json:"collaborators"`
	GrcContacts             []UserOrGroup       `json:"service_owners"`
	SecurityPolicies        []SecurityPolicy    `json:"security_policies"`
	Projects                []Project           `json:"projects"`
	DocumentationUrl        string              `json:"documentation_url"`
	Risks                   []Risk              `json:"risks"`
	ThirdPartyRisks         []ThirdPartyRisk    `json:"third_party_risks"`
	Classifications         []Tag               `json:"classifications"`
}

func (p *SecurityService) GetId() int32 {
	return p.Id
}

func (p *SecurityService) Link(base string) string {
	return ErambaViewLink(base, "security-services", p.Id)
}

func (p *SecurityService) MarshalJSON() ([]byte, error) {
	type Alias SecurityService
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, RiskExceptionSkippedFields)
}
