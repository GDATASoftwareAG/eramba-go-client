package model

import (
	"encoding/json"
	"fmt"
	"maps"
	"strings"
)

const minRiskClassificationKeySplits = 3

func UnmarshalRiskClassification(
	prefix string, data []byte,
) (analysis, treatment map[int32]RiskClassification, err error) {
	analysis = map[int32]RiskClassification{}
	treatment = map[int32]RiskClassification{}
	baseFields := map[string]any{}
	if err := json.Unmarshal(data, &baseFields); err != nil {
		return analysis, treatment, err
	}
	keys := maps.Keys(baseFields)
	for key := range keys {
		if !strings.HasPrefix(key, prefix) {
			delete(baseFields, key)
			continue
		}
		data, err := json.Marshal(baseFields[key])
		if err != nil {
			return analysis, treatment, err
		}
		riskClassification := RiskClassification{}
		if err := json.Unmarshal(data, &riskClassification); err != nil {
			return analysis, treatment, err
		}
		id, typeId, err := convertToTypeId(prefix, key)
		if err != nil {
			return analysis, treatment, err
		}
		if id == "1" {
			treatment[typeId] = riskClassification
		}
		if id == "0" {
			analysis[typeId] = riskClassification
		}
	}
	return analysis, treatment, nil
}

func UnmarshalWithCustomFieldsAndRiskClassification[T any](
	data []byte, alias *T, prefix string,
) (customFields CustomFields, analysis, treatment map[int32]RiskClassification, err error) {
	customFields, err = UnmarshalWithCustomFields(data, alias)
	if err != nil {
		return nil, nil, nil, err
	}
	analysis, treatment, err = UnmarshalRiskClassification(prefix, data)
	if err != nil {
		return nil, nil, nil, err
	}
	return customFields, analysis, treatment, nil
}

func MarshalRiskClassification(
	s string,
	classificationsAnalysis, classificationsTreatment map[int32]RiskClassification,
) map[string]any {
	fields := map[string]any{}
	for k, v := range classificationsAnalysis {
		fields[fmt.Sprintf("%s_0__type_%d", s, k)] = v
	}
	for k, v := range classificationsTreatment {
		fields[fmt.Sprintf("%s_1__type_%d", s, k)] = v
	}
	return fields
}
