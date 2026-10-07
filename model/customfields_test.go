package model_test

import (
	"encoding/json"
	"testing"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func Test_CustomFields_MarshalUnmarshal_Int(t *testing.T) {
	tp := &model.ThirdParty{
		Id:    1,
		Title: "Acme",
	}
	tp.CustomFields = model.CustomFields{}
	tp.CustomFields.SetInt("custom_field_1", 42)

	data, err := json.Marshal(tp)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	got := &model.ThirdParty{}
	if err := json.Unmarshal(data, got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	value, ok := got.CustomFields.GetInt("custom_field_1")
	if !ok {
		t.Fatalf("GetInt() ok = false, want true (data: %s)", data)
	}
	if value != 42 {
		t.Errorf("GetInt() = %d, want 42", value)
	}
}

func Test_CustomFields_MarshalUnmarshal_String(t *testing.T) {
	tp := &model.ThirdParty{
		Id:    1,
		Title: "Acme",
	}
	tp.CustomFields = model.CustomFields{}
	tp.CustomFields.SetString("custom_field_1", "hello world")

	data, err := json.Marshal(tp)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	got := &model.ThirdParty{}
	if err := json.Unmarshal(data, got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	value, ok := got.CustomFields.GetString("custom_field_1")
	if !ok {
		t.Fatalf("GetString() ok = false, want true (data: %s)", data)
	}
	if value != "hello world" {
		t.Errorf("GetString() = %q, want %q", value, "hello world")
	}
}

func Test_CustomFields_MarshalUnmarshal_StringSlices(t *testing.T) {
	tp := &model.ThirdParty{
		Id:    1,
		Title: "Acme",
	}
	tp.CustomFields = model.CustomFields{}
	want := []string{"one", "two", "three"}
	tp.CustomFields.SetStringSlices("custom_field_1", want)

	data, err := json.Marshal(tp)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	got := &model.ThirdParty{}
	if err := json.Unmarshal(data, got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	value, ok := got.CustomFields.GetStringSlices("custom_field_1")
	if !ok {
		t.Fatalf("GetStringSlices() ok = false, want true (data: %s)", data)
	}
	if len(value) != len(want) {
		t.Fatalf("GetStringSlices() = %v, want %v", value, want)
	}
	for i := range want {
		if value[i] != want[i] {
			t.Errorf("GetStringSlices()[%d] = %q, want %q", i, value[i], want[i])
		}
	}
}

func Test_CustomFields_Get_MissingKey(t *testing.T) {
	cf := model.CustomFields{}

	if value, ok := cf.GetInt("custom_field_missing"); ok || value != 0 {
		t.Errorf("GetInt() = (%d, %v), want (0, false)", value, ok)
	}
	if value, ok := cf.GetString("custom_field_missing"); ok || value != "" {
		t.Errorf("GetString() = (%q, %v), want (\"\", false)", value, ok)
	}
	if value, ok := cf.GetStringSlices("custom_field_missing"); ok || len(value) != 0 {
		t.Errorf("GetStringSlices() = (%v, %v), want ([], false)", value, ok)
	}
}

func Test_CustomFields_Get_WrongType(t *testing.T) {
	tp := &model.ThirdParty{
		Id:    1,
		Title: "Acme",
	}
	tp.CustomFields = model.CustomFields{}
	tp.CustomFields.SetString("custom_field_1", "hello world")

	data, err := json.Marshal(tp)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	got := &model.ThirdParty{}
	if err := json.Unmarshal(data, got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}

	if value, ok := got.CustomFields.GetInt("custom_field_1"); ok || value != 0 {
		t.Errorf("GetInt() on a string field = (%d, %v), want (0, false)", value, ok)
	}
	if value, ok := got.CustomFields.GetStringSlices("custom_field_1"); ok || len(value) != 0 {
		t.Errorf("GetStringSlices() on a string field = (%v, %v), want ([], false)", value, ok)
	}
}
