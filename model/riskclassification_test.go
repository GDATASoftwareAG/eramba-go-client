package model_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/gdatasoftwareag/eramba-go-client/model"
)

func Test_MarshalRiskClassification(t *testing.T) {
	analysis := map[int32]model.RiskClassification{
		1: {RiskClassificationId: 10},
	}
	treatment := map[int32]model.RiskClassification{
		2: {RiskClassificationId: 20},
	}

	fields := model.MarshalRiskClassification("risk_classifications__risks", analysis, treatment)

	data, err := json.Marshal(fields)
	if err != nil {
		t.Fatalf("Marshal() error = %v", err)
	}

	want := map[string]any{
		"risk_classifications__risks_0__type_1": float64(10),
		"risk_classifications__risks_1__type_2": float64(20),
	}
	got := map[string]any{}
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("Unmarshal() error = %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MarshalRiskClassification() = %v, want %v", got, want)
	}
}

func Test_UnmarshalRiskClassification(t *testing.T) {
	data := []byte(`{
		"risk_classifications__risks_0__type_1": {"risk_classification_id": 10},
		"risk_classifications__risks_1__type_2": {"risk_classification_id": 20},
		"unrelated_field": "ignored"
	}`)

	analysis, treatment, err := model.UnmarshalRiskClassification("risk_classifications__risks", data)
	if err != nil {
		t.Fatalf("UnmarshalRiskClassification() error = %v", err)
	}

	wantAnalysis := map[int32]model.RiskClassification{1: {RiskClassificationId: 10}}
	wantTreatment := map[int32]model.RiskClassification{2: {RiskClassificationId: 20}}
	if !reflect.DeepEqual(analysis, wantAnalysis) {
		t.Errorf("analysis = %+v, want %+v", analysis, wantAnalysis)
	}
	if !reflect.DeepEqual(treatment, wantTreatment) {
		t.Errorf("treatment = %+v, want %+v", treatment, wantTreatment)
	}
}
