package model

type Comment struct {
	Id         int32  `json:"id"`
	Message    string `json:"message"`
	ForeignKey int32  `json:"foreign_key"`
	UserId     int32  `json:"user_id"`
}

func (p *Comment) MarshalJSON() ([]byte, error) {
	type Alias Comment
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, []string{"id", "foreign_key", "user_id"})
}
