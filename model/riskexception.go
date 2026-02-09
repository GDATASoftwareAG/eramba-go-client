package model

var RiskExceptionSkippedFields = []string{
	"id",
}

type RiskException struct {
	Id          int32         `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Expiration  *ErambaDate   `json:"expiration"`
	ClosureDate *ErambaDate   `json:"closure_date"`
	Tags        []Tag         `json:"tags"`
	Requesters  []UserOrGroup `json:"requesters"`
	GrcContacts []UserOrGroup `json:"owners"`
	Status      int32         `json:"status"`
}

func (p *RiskException) GetId() int32 {
	return p.Id
}

func (p *RiskException) Link(base string) string {
	return ErambaViewLink(base, "risk-exceptions", p.Id)
}

func (p *RiskException) MarshalJSON() ([]byte, error) {
	type Alias RiskException
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, RiskExceptionSkippedFields)
}
