package model

import "fmt"

type User struct {
	Id      int32   `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Email   string  `json:"email"`
	Groups  []Group `json:"groups,omitempty"`
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

func (p *User) MarshalJSON() ([]byte, error) {
	type Alias User
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, []string{})
}
