package model

import "encoding/json"

type ThirdParties []*ThirdParty

func (p ThirdParties) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type ThirdParty struct {
	Id               int32  `json:"id"`
	ThirdPartyTypeId int32  `json:"third_party_type_id"`
	Title            string `json:"name"`
	Description      string `json:"description"`

	Sponsors      []UserOrGroup `json:"sponsors"`
	GrcContacts   []UserOrGroup `json:"grc_contacts"`
	BusinessUnits []*OnlyId     `json:"business_units"`
	Processes     []*OnlyId     `json:"processes"`
	Legals        []*OnlyId     `json:"legals"`

	CustomFields map[string]CustomField `json:"-"`
}

func (p *ThirdParty) GetId() int32 {
	return p.Id
}

func (p *ThirdParty) Link(base string) string {
	return ErambaViewLink(base, "third-parties", p.Id)
}

func (p *ThirdParty) UnmarshalJSON(data []byte) error {
	type Alias ThirdParty // avoid recursion
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
	return nil
}

func (p *ThirdParty) MarshalJSON() ([]byte, error) {
	type Alias ThirdParty
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, p.CustomFields, []string{})
}
