package model

import "encoding/json"

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
	return ErambaViewLink(base, "third-parties", p.Id)
}
