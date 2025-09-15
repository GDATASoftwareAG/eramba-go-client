package model

import (
	"encoding/json"
	"maps"
	"strings"
)

type ProjectStatus int32

const (
	ProjectStatusDone    ProjectStatus = 3
	ProjectStatusOngoing ProjectStatus = 2
	ProjectStatusPlanned ProjectStatus = 1
)

var ProjectSkippedFields = []string{
	"id",
	"risks",
	"third_party_risks",
}

type Projects []*Project

func (p Projects) MarshalJSON() ([]byte, error) {
	list := ExtractPatchListId(p)
	return json.Marshal(list)
}

type Project struct {
	Id              int32                  `json:"id"`
	Title           string                 `json:"title"`
	Description     string                 `json:"goal"`
	PlanBudget      string                 `json:"plan_budget,omitempty"`
	Start           *ErambaDate            `json:"start"`
	Deadline        *ErambaDate            `json:"deadline"`
	ProjectStatus   ProjectStatus          `json:"project_status_id"`
	Tags            []Tag                  `json:"tags"`
	Owners          []UserOrGroup          `json:"owners"`
	Contacts        []UserOrGroup          `json:"contacts"`
	Risks           []Risk                 `json:"risks"`
	ThirdPartyRisks []ThirdPartyRisk       `json:"third_party_risks"`
	CustomFields    map[string]CustomField `json:"-"`
}

func (p *Project) UnmarshalJSON(data []byte) error {
	type Alias Project // avoid recursion
	aux := &struct {
		*Alias
		// Capture all unknown fields
		CustomFields map[string]interface{} `json:"-"`
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux.CustomFields); err != nil {
		return err
	}
	if err := json.Unmarshal(data, &aux.Alias); err != nil {
		return err
	}
	customFields := map[string]CustomField{}
	keys := maps.Keys(aux.CustomFields)
	for key := range keys {
		if !strings.HasPrefix(key, "custom_field_") {
			delete(aux.CustomFields, key)
		} else {
			data, err := json.Marshal(aux.CustomFields[key])
			if err != nil {
				return err
			}
			customField := CustomField{}
			if err := json.Unmarshal(data, &customField); err != nil {
				return err
			}
			customFields[key] = customField
		}
	}
	p.CustomFields = customFields
	return nil
}

func (p *Project) MarshalJSON() ([]byte, error) {
	type Alias Project
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, p.CustomFields, ReviewSkippedFields)
}

func (p *Project) GetId() int32 {
	return p.Id
}

func (p *Project) Link(base string) string {
	return buildLink(base, "projects", p.Id)
}
