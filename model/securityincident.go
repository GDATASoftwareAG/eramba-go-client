package model

import (
	"encoding/json"
	"fmt"
)

type SecurityIncidentType string

const (
	SecurityIncidentTypeEvent            SecurityIncidentType = "event"
	SecurityIncidentTypePossibleIncident SecurityIncidentType = "possible-incident"
	SecurityIncidentTypeIncident         SecurityIncidentType = "incident"
)

type SecurityIncidentStatus int32

const (
	SecurityIncidentStatusOpen   SecurityIncidentStatus = 2
	SecurityIncidentStatusClosed SecurityIncidentStatus = 3
)

type SecurityIncidentStageStatus int32

const (
	SecurityIncidentStageStatusIncomplete SecurityIncidentStageStatus = 0
	SecurityIncidentStageStatusCompleted  SecurityIncidentStageStatus = 1
)

// Fields returned by the API that are read-only and must not be sent on create/update.
var SecurityIncidentSkippedFields = []string{
	FieldId,
	"assets",
	"third_parties",
	"security_incident_stages",
	"security_incident_stages_security_incidents",
	"created",
	"edited",
	"modified",
}

var _ ErambaType = (*SecurityIncident)(nil)

type SecurityIncident struct {
	Id          int32                  `json:"id"`
	Title       string                 `json:"title"`
	Description string                 `json:"description"`
	Type        SecurityIncidentType   `json:"type"`
	Status      SecurityIncidentStatus `json:"security_incident_status_id"`
	Owners      UsersAndGroups         `json:"owners"`
	Contacts    UsersAndGroups         `json:"contacts"`
	OpenDate    *ErambaDate            `json:"open_date,omitempty"`
	ClosureDate *ErambaDate            `json:"closure_date"`
	Reporter    string                 `json:"reporter"`
	Victim      string                 `json:"victim"`

	Classifications      Classifications      `json:"classifications"`
	Risks                Risks                `json:"risks"`
	ThirdPartyRisks      ThirdPartyRisks      `json:"third_party_risks"`
	BusinessContinuities BusinessContinuities `json:"business_continuities"`
	SecurityServices     SecurityServices     `json:"security_services"`

	// read-only fields
	Assets         Assets                          `json:"assets"`
	ThirdParties   ThirdParties                    `json:"third_parties"`
	Stages         []SecurityIncidentStage         `json:"security_incident_stages"`
	StageInstances []SecurityIncidentStageInstance `json:"security_incident_stages_security_incidents"`
	Created        string                          `json:"created,omitempty"`
	Edited         string                          `json:"edited,omitempty"`
	Modified       string                          `json:"modified,omitempty"`

	CustomFields CustomFields `json:"-"`
}

func (p *SecurityIncident) GetId() int32 {
	return p.Id
}

func (p *SecurityIncident) Link(base string) string {
	return ErambaViewLink(base, "security-incidents", p.Id)
}

func (p *SecurityIncident) UnmarshalJSON(data []byte) error {
	type Alias SecurityIncident // avoid recursion
	customFields, err := UnmarshalWithCustomFields(data, (*Alias)(p))
	if err != nil {
		return err
	}
	p.CustomFields = customFields
	return nil
}

func (p *SecurityIncident) MarshalJSON() ([]byte, error) {
	type Alias SecurityIncident
	return MarshalWithSpecialFields(Alias(*p), p.CustomFields, map[string]any{}, SecurityIncidentSkippedFields)
}

type SecurityIncidentStage struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SecurityIncidentStageInstance struct {
	Id                      int32                       `json:"id"`
	SecurityIncidentId      int32                       `json:"security_incident_id"`
	SecurityIncidentStageId int32                       `json:"security_incident_stage_id"`
	Status                  SecurityIncidentStageStatus `json:"status"`
	StageName               string                      `json:"stage_name"`
	StageDescription        string                      `json:"stage_description"`
}

// Classifications are sent as a list of strings. The API may return them
// either as strings or as objects with a "title"/"name" field.
type Classifications []string

func (c Classifications) MarshalJSON() ([]byte, error) {
	if c == nil {
		return []byte("[]"), nil
	}
	return json.Marshal([]string(c))
}

func (c *Classifications) UnmarshalJSON(data []byte) error {
	raw := []json.RawMessage{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	out := make(Classifications, 0, len(raw))
	for _, item := range raw {
		var s string
		if err := json.Unmarshal(item, &s); err == nil {
			if s != "" {
				out = append(out, s)
			}
			continue
		}
		obj := struct {
			Title string `json:"title"`
			Name  string `json:"name"`
		}{}
		if err := json.Unmarshal(item, &obj); err != nil {
			return err
		}
		switch {
		case obj.Title != "":
			out = append(out, obj.Title)
		case obj.Name != "":
			out = append(out, obj.Name)
		default:
			return fmt.Errorf("unsupported classification: %s", item)
		}
	}
	*c = out
	return nil
}
