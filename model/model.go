package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

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
