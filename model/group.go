package model

import (
	"encoding/json"
	"fmt"
)

var GroupSkippedFields = []string{
	FieldId,
	"users",
}
var _ ErambaType = (*Group)(nil)

type Group struct {
	Id          int32  `json:"id"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Users       []User `json:"users,omitempty"`
}

func (p *Group) GetId() int32 {
	return p.Id
}

func (p *Group) Link(base string) string {
	return ErambaViewLink(base, "groups", p.Id)
}

func (p *Group) GenerateUserOrGroup() UserOrGroup {
	return UserOrGroup{
		ObjectKey: fmt.Sprintf("Group-%d", p.Id),
		Group: struct {
			Name string `json:"name"`
		}{
			Name: p.Name,
		},
	}
}

func (p *Group) MarshalJSON() ([]byte, error) {
	type Alias Group
	return MarshalWithSkippingFields(Alias(*p), GroupSkippedFields)
}

type Groups []*Group

func (p Groups) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}
