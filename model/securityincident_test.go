package model_test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

const securityIncidentResponse = `{
	"id": 7,
	"title": "Phishing",
	"description": "Mail campaign",
	"type": "possible-incident",
	"security_incident_status_id": 2,
	"owners": {"users": [{"id": 1, "name": "Admin"}], "groups": [{"id": 3, "name": "SOC"}]},
	"contacts": {"users": [], "groups": [{"id": 4, "name": "IT"}]},
	"open_date": "2026-01-15",
	"closure_date": null,
	"reporter": "Alice",
	"victim": "Bob",
	"classifications": ["phishing", {"title": "mail"}],
	"security_incident_stages_security_incidents": [
		{"id": 1, "security_incident_id": 7, "security_incident_stage_id": 2, "status": 1, "stage_name": "Triage"}
	],
	"security_incident_stages": [{"id": 2, "name": "Triage"}],
	"risks": [{"id": 11, "title": "r"}],
	"third_party_risks": [],
	"business_continuities": [{"id": 12}],
	"assets": [{"id": 13}],
	"third_parties": [{"id": 14}],
	"security_services": [{"id": 15}],
	"created": "2026-01-15 10:00:00",
	"edited": null,
	"modified": "2026-01-16 10:00:00",
	"custom_field_severity": {"value": "high", "custom_field_id": 5}
}`

func userAndGroupNames(u model.UsersAndGroups) []string {
	names := make([]string, 0, len(u.Users)+len(u.Groups))
	for _, user := range u.Users {
		names = append(names, fmt.Sprintf("User-%d %s", user.Id, user.Name))
	}
	for _, group := range u.Groups {
		names = append(names, fmt.Sprintf("Group-%d %s", group.Id, group.Name))
	}
	return names
}

func Test_SecurityIncident_Unmarshal(t *testing.T) {
	incident := model.SecurityIncident{}
	if err := json.Unmarshal([]byte(securityIncidentResponse), &incident); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if incident.Id != 7 || incident.Type != model.SecurityIncidentTypePossibleIncident ||
		incident.Status != model.SecurityIncidentStatusOpen {
		t.Errorf("unexpected base fields: %+v", incident)
	}
	if got := userAndGroupNames(incident.Owners); !reflect.DeepEqual(got, []string{"User-1 Admin", "Group-3 SOC"}) {
		t.Errorf("owners = %v", got)
	}
	if got := incident.Contacts.ObjectKeys(); !reflect.DeepEqual(got, []string{"Group-4"}) {
		t.Errorf("contacts = %v", got)
	}
	if incident.OpenDate.String() != "2026-01-15" || incident.ClosureDate != nil {
		t.Errorf("dates = %v / %v", incident.OpenDate, incident.ClosureDate)
	}
	if !reflect.DeepEqual(incident.Classifications, model.Classifications{"phishing", "mail"}) {
		t.Errorf("classifications = %v", incident.Classifications)
	}
	if len(incident.StageInstances) != 1 ||
		incident.StageInstances[0].Status != model.SecurityIncidentStageStatusCompleted {
		t.Errorf("stage instances = %+v", incident.StageInstances)
	}
	if v, ok := incident.CustomFields.GetString("custom_field_severity"); !ok || v != "high" {
		t.Errorf("custom field = %v, %v", v, ok)
	}
}

func Test_SecurityIncident_Marshal(t *testing.T) {
	incident := model.SecurityIncident{}
	if err := json.Unmarshal([]byte(securityIncidentResponse), &incident); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	data, err := json.Marshal(&incident)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	got := map[string]any{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal map failed: %v", err)
	}
	want := map[string]any{
		"title":                       "Phishing",
		"description":                 "Mail campaign",
		"type":                        "possible-incident",
		"security_incident_status_id": float64(2),
		"owners":                      []any{"User-1", "Group-3"},
		"contacts":                    []any{"Group-4"},
		"open_date":                   "2026-01-15",
		"reporter":                    "Alice",
		"victim":                      "Bob",
		"classifications":             []any{"phishing", "mail"},
		"closure_date":                nil,
		"risks":                       []any{"11"},
		"third_party_risks":           []any{},
		"business_continuities":       []any{"12"},
		"security_services":           []any{"15"},
		"custom_field_severity":       map[string]any{"value": "high", "custom_field_id": float64(5)},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("marshal mismatch\n got: %v\nwant: %v", got, want)
	}
}

func Test_SecurityIncident_MarshalNew(t *testing.T) {
	incident := model.SecurityIncident{
		Title:  "New",
		Type:   model.SecurityIncidentTypeEvent,
		Status: model.SecurityIncidentStatusOpen,
	}
	incident.Owners.Users = append(incident.Owners.Users, &model.User{Id: 1})
	incident.Contacts.Groups = append(incident.Contacts.Groups, &model.Group{Id: 2})
	data, err := json.Marshal(&incident)
	if err != nil {
		t.Fatalf("marshal failed: %v", err)
	}
	got := map[string]any{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal map failed: %v", err)
	}
	if !reflect.DeepEqual(got["owners"], []any{"User-1"}) || !reflect.DeepEqual(got["contacts"], []any{"Group-2"}) {
		t.Errorf("owners/contacts = %v / %v", got["owners"], got["contacts"])
	}
	for _, k := range []string{"id", "open_date"} {
		if _, ok := got[k]; ok {
			t.Errorf("unexpected key %q in %s", k, data)
		}
	}
	// cleared values must be sent explicitly so Put can remove them
	if v, ok := got["closure_date"]; !ok || v != nil {
		t.Errorf("closure_date = %v (present %v), want null", v, ok)
	}
	for _, k := range []string{"risks", "third_party_risks", "business_continuities", "security_services", "classifications"} {
		if !reflect.DeepEqual(got[k], []any{}) {
			t.Errorf("%s = %v, want []", k, got[k])
		}
	}
}

func Test_UsersAndGroups_Unmarshal(t *testing.T) {
	tests := []struct {
		name string
		data string
		want []string
	}{
		{name: "object", data: `{"users":[{"id":1}],"groups":[{"id":2}]}`, want: []string{"User-1", "Group-2"}},
		{name: "empty list", data: `[]`, want: []string{}},
		{name: "null", data: `null`, want: []string{}},
		{name: "object keys", data: `["User-5","Group-6"]`, want: []string{"User-5", "Group-6"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := model.UsersAndGroups{}
			if err := json.Unmarshal([]byte(tt.data), &got); err != nil {
				t.Fatalf("unmarshal failed: %v", err)
			}
			if keys := got.ObjectKeys(); !reflect.DeepEqual(keys, tt.want) {
				t.Errorf("ObjectKeys() = %v, want %v", keys, tt.want)
			}
		})
	}
	invalid := model.UsersAndGroups{}
	if err := json.Unmarshal([]byte(`["Team-1"]`), &invalid); err == nil {
		t.Error("expected error for invalid reference")
	}
}

func Test_Classifications_Unmarshal(t *testing.T) {
	got := model.Classifications{}
	if err := json.Unmarshal([]byte(`["a", "", {"title":"b"}, {"name":"c"}]`), &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if !reflect.DeepEqual(got, model.Classifications{"a", "b", "c"}) {
		t.Errorf("classifications = %v", got)
	}
	if err := json.Unmarshal([]byte(`[{"id":5}]`), &got); err == nil {
		t.Error("expected error for classification without title or name")
	}
}
