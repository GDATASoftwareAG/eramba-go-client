package model

import "fmt"

var _ ErambaType = (*User)(nil)

type User struct {
	Id      int32   `json:"id"`
	Name    string  `json:"name"`
	Surname string  `json:"surname"`
	Email   string  `json:"email"`
	Groups  []Group `json:"groups,omitempty"`
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

func (p *User) MarshalJSON() ([]byte, error) {
	type Alias User
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, []string{})
}
