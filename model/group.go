package model

import "fmt"

type Group struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Slug  string `json:"slug"`
	Users []User `json:"users,omitempty"`
}

func (p *Group) GetId() int32 {
	return p.Id
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
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, []string{})
}
