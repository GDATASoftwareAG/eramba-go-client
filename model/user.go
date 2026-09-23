package model

import (
	"fmt"
)

var _ ErambaType = (*User)(nil)

type User struct {
	Id      int32   `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Email   string  `json:"email"`
	Groups  []Group `json:"groups,omitempty"`

	CustomFields CustomFields `json:"-"`
}

func (p *User) Link(base string) string {
	return ErambaViewLink(base, "users", p.Id)
}

func (p *User) GetId() int32 {
	return p.Id
}

func (p *User) GenerateUserOrGroup() UserOrGroup {
	return UserOrGroup{
		ObjectKey: fmt.Sprintf("User-%d", p.Id),
		User: struct {
			Name string `json:"name"`
		}{
			Name: p.Name,
		},
	}
}

func (p *User) UnmarshalJSON(data []byte) error {
	type Alias User // avoid recursion
	customFields, err := UnmarshalWithCustomFields(data, (*Alias)(p))
	if err != nil {
		return err
	}
	p.CustomFields = customFields
	return nil
}

func (p *User) MarshalJSON() ([]byte, error) {
	type Alias User
	return MarshalWithSpecialFields(Alias(*p), p.CustomFields, map[string]any{}, []string{})
}
