package model

import (
	"fmt"
)

var UserSkippedFields = []string{
	FieldId,
}

var _ ErambaType = (*User)(nil)

type User struct {
	Id      int32  `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Login   string `json:"login"`
	Status  int    `json:"status"`
	Groups  Groups `json:"groups,omitempty"`

	LocalAccount            bool `json:"local_account"`
	ApiAllow                bool `json:"api_allow"`
	MainPortal              bool `json:"main_portal"`
	VendorAssessmentsPortal bool `json:"vendor_assessments_portal"`
	AccountReviewsPortal    bool `json:"account_reviews_portal"`
	AwarenessPortal         bool `json:"awareness_portal"`
	PolicyPortal            bool `json:"policy_portal"`

	UserTemplateId int `json:"user_template_id"`

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
	return MarshalWithSpecialFields(Alias(*p), p.CustomFields, map[string]any{}, UserSkippedFields)
}
