package model

import (
	"encoding/json"
)

type BusinessUnit struct {
	Id int32 `json:"id"`
}

func (b BusinessUnit) GetId() int32 {
	return b.Id
}

func (b BusinessUnit) Link(base string) string {
	return buildLink(base, "business-units", b.Id)
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

type ThirdParties []*ThirdParty

func (p ThirdParties) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type ThirdParty struct {
	Id          int32  `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
}

func (p *ThirdParty) GetId() int32 {
	return p.Id
}

func (p *ThirdParty) Link(base string) string {
	return buildLink(base, "third-parties", p.Id)
}

type RiskClassification struct {
	RiskClassificationId int32 `json:"risk_classification_id"`
}

func (o RiskClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.RiskClassificationId)
}

type Comment struct {
	Message string `json:"message"`
}
