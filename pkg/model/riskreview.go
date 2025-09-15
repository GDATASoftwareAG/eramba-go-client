package model

import (
	"time"
)

var ReviewSkippedFields = []string{
	"id",
	"model",
	"third_party_risks",
}

type Review struct {
	Id          int32         `json:"id"`
	Model       string        `json:"model"`
	ForeignKey  int32         `json:"foreign_key"`
	Completed   int32         `json:"completed"`
	Description string        `json:"description"`
	PlannedDate *ErambaDate   `json:"planned_date"`
	ActualDate  *ErambaDate   `json:"actual_date"`
	Reviewers   []UserOrGroup `json:"reviewers"`
}

func (p *Review) After(review *Review) bool {
	if review == nil {
		return true
	}
	return time.Time(*p.ActualDate).After(time.Time(*review.ActualDate))
}

func (p *Review) MarshalJSON() ([]byte, error) {
	type Alias Review
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, ReviewSkippedFields)
}
