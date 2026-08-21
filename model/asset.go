package model

import (
	"encoding/json"
	"fmt"
)

type AssetClassification struct {
	Id                      int32  `json:"id"`
	Value                   int32  `json:"value"`
	Name                    string `json:"name"`
	TypeId                  int32  `json:"asset_classification_type_id"`
	AssetClassificationType struct {
		Name string `json:"name"`
	} `json:"asset_classification_type,omitempty"`
}

type AssetLabelId int32

func (p *AssetLabelId) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte(""), nil
	}
	return fmt.Appendf([]byte{}, "%d", p), nil
}

type Risks []*Risk

func (p Risks) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type Assets []*Asset

func (p Assets) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type Asset struct {
	Id                   int32                 `json:"id"`
	Title                string                `json:"name"`
	BusinessUnits        []*OnlyId             `json:"business_units"`
	AssetOwners          []UserOrGroup         `json:"asset_owners"`
	AssetUsers           []UserOrGroup         `json:"asset_users"`
	GrcContacts          []UserOrGroup         `json:"grc_contacts"`
	AssetGuardians       []UserOrGroup         `json:"asset_guardians"`
	AssetClassifications []AssetClassification `json:"asset_classifications"`
	RelatedAssets        Assets                `json:"related_assets"`
	AssetLabelId         *AssetLabelId         `json:"asset_label_id"`
	Legals               []*OnlyId             `json:"legals"`
	Description          string                `json:"description"`
	AssetMediaTypeId     int32                 `json:"asset_media_type_id"`
	Risks                Risks                 `json:"risks"`
}

func (p *Asset) GetId() int32 {
	return p.Id
}

func (p *Asset) Link(base string) string {
	return ErambaViewLink(base, "assets", p.Id)
}

var AssetSkippedFields = []string{
	FieldId,
	"asset_classifications",
}

func (p *Asset) MarshalJSON() ([]byte, error) {
	extraFields := make(map[string]any)
	for _, classification := range p.AssetClassifications {
		extraFields[fmt.Sprintf("asset_classifications_%d", classification.TypeId)] = []int32{classification.Id}
	}
	return MarshalWithSpecialFields(p, CustomFields{}, extraFields, AssetSkippedFields)
}
