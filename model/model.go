package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"slices"
	"strconv"
	"strings"
)

const minRiskClassificationKeySplits = 3

type BusinessUnit struct {
	Id int32 `json:"id"`
}

func (b BusinessUnit) GetId() int32 {
	return b.Id
}

func (b BusinessUnit) Link(base string) string {
	return ErambaViewLink(base, "business-units", b.Id)
}

type UserOrGroup struct {
	ObjectKey string `json:"object_key"`
	Group     struct {
		Name string `json:"name"`
	} `json:"group,omitempty"`
	User struct {
		Name string `json:"name"`
	} `json:"user,omitempty"`
}

func (o UserOrGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ObjectKey)
}

type RiskAppetiteThreshold struct {
	Id    int32  `json:"risk_appetite_threshold_id"`
	Title string `json:"title,omitempty"`
}

type RiskThreat struct {
	Id              int32    `json:"id"`
	Title           string   `json:"name"`
	AssetMediaTypes []OnlyId `json:"asset_media_types"`
}

type Tag struct {
	Title string `json:"title"`
}

func (o Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Title)
}

type RiskScore struct {
	Score int32 `json:"score"`
}

type RiskClassification struct {
	RiskClassificationId int32 `json:"risk_classification_id"`
}

func (o RiskClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.RiskClassificationId)
}

func convertToTypeId(prefix, key string) (id string, typeId int32, err error) {
	key = strings.Replace(key, prefix, "", 1)

	splits := slices.DeleteFunc(strings.Split(key, "_"), func(s string) bool {
		return s == ""
	})
	if len(splits) < minRiskClassificationKeySplits {
		return "", 0, errors.New("not enough")
	}
	firstElement := splits[0]
	lastElement := splits[len(splits)-1]
	num, err := strconv.ParseInt(lastElement, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return firstElement, 0, err
	}

	return firstElement, int32(num), nil
}

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
