package model

import (
	"encoding/json"
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
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type Project struct {
	Id              int32                  `json:"id"`
	Title           string                 `json:"title"`
	Description     string                 `json:"goal"`
	PlanBudget      *int32                 `json:"plan_budget"`
	Start           *ErambaDate            `json:"start"`
	Deadline        *ErambaDate            `json:"deadline"`
	ProjectStatus   ProjectStatus          `json:"project_status_id"`
	Tags            []Tag                  `json:"tags"`
	Owners          []UserOrGroup          `json:"owners"`
	GrcContacts     []UserOrGroup          `json:"contacts"`
	Risks           []Risk                 `json:"risks"`
	ThirdPartyRisks []ThirdPartyRisk       `json:"third_party_risks"`
	CustomFields    map[string]CustomField `json:"-"`
}

func (p *Project) UnmarshalJSON(data []byte) error {
	type Alias Project // avoid recursion
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

func (p *Project) MarshalJSON() ([]byte, error) {
	type Alias Project
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, p.CustomFields, ProjectSkippedFields)
}

func (p *Project) GetId() int32 {
	return p.Id
}

func (p *Project) Link(base string) string {
	return ErambaViewLink(base, "projects", p.Id)
}
