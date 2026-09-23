package model

import "encoding/json"

var RiskExceptionSkippedFields = []string{
	FieldId,
	FieldRisks,
	FieldThirdPartyRisks,
	FieldBusinessContinuities,
}

type RiskExceptions []*RiskException

func (p RiskExceptions) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type RiskException struct {
	Id                   int32                `json:"id"`
	Title                string               `json:"title"`
	Description          string               `json:"description"`
	Expiration           *ErambaDate          `json:"expiration"`
	ClosureDate          *ErambaDate          `json:"closure_date"`
	StartDate            *ErambaDate          `json:"start_date"`
	Tags                 []Tag                `json:"tags"`
	Requesters           []UserOrGroup        `json:"requesters"`
	GrcContacts          []UserOrGroup        `json:"owners"`
	Status               int32                `json:"status"`
	Risks                []Risk               `json:"risks"`
	ThirdPartyRisks      []ThirdPartyRisk     `json:"third_party_risks"`
	BusinessContinuities []BusinessContinuity `json:"business_continuities"`
}

func (p *RiskException) GetId() int32 {
	return p.Id
}

func (p *RiskException) Link(base string) string {
	return ErambaViewLink(base, "risk-exceptions", p.Id)
}

func (p *RiskException) MarshalJSON() ([]byte, error) {
	type Alias RiskException
	return MarshalWithSkippingFields(Alias(*p), RiskExceptionSkippedFields)
}
