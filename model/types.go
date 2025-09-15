package model

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type CustomField struct {
	Value         interface{} `json:"value"`
	CustomFieldId int32       `json:"custom_field_id"`
}

func MarshalWithSkippingFields[T any](
	p T,
	customFields map[string]CustomField,
	skippedFields []string,
) ([]byte, error) {
	data, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	// Turn into map for merging
	out := make(map[string]interface{})
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	for _, k := range skippedFields {
		delete(out, k)
	}
	// Add extra fields back
	for k, v := range customFields {
		out[k] = v
	}

	return json.Marshal(out)
}

func extractPatchListId[K ErambaType](erambaObjects []K) []string {
	objects := map[string]bool{}
	if len(erambaObjects) == 0 {
		return []string{}
	}
	for _, erambaObject := range erambaObjects {
		objects[strconv.Itoa(int(erambaObject.GetId()))] = true
	}
	keys := make([]string, len(objects))
	i := 0
	for k := range objects {
		keys[i] = k
		i++
	}
	return keys
}

type ErambaType interface {
	GetId() int32
	Link(base string) string
}

type OnlyId struct {
	Id int32 `json:"id"`
}

func (p *OnlyId) GetId() int32 {
	return p.Id
}

func (p *OnlyId) Link(_ string) string {
	return ""
}

func (p *OnlyId) MarshalJSON() ([]byte, error) {
	if p == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf("%d", p.Id)), nil
}

const DateFormat = "2006-01-02"

type ErambaDate time.Time

func RefErambaDate(date time.Time) *ErambaDate {
	erambaDate := ErambaDate(date)
	return &erambaDate
}

func (t *ErambaDate) UnmarshalJSON(b []byte) (err error) {
	if string(b) == "null" {
		return nil
	}

	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	*t = ErambaDate(date)
	return
}

func (t *ErambaDate) MarshalJSON() ([]byte, error) {
	if t == nil {
		return []byte("null"), nil
	}
	return []byte(time.Time(*t).Format(`"2006-01-02"`)), nil
}

func (t *ErambaDate) String() string {
	if t == nil {
		return ""
	}
	return time.Time(*t).Format(DateFormat)
}

func (t *ErambaDate) IsEqual(b time.Time) bool {
	return b.Format(DateFormat) == t.String()
}

func buildLink(base, tool string, id int32) string {
	filter := "filter%5Bid%5D%5Boperator%5D=%24in&filter%5Bid%5D%5Bvalue%5D%5B0%5D="
	return fmt.Sprintf("%s/%s?%s%d", base, tool, filter, id)
}
