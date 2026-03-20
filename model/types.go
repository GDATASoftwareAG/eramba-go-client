package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"maps"
	"strconv"
	"strings"
	"time"
	"unicode"
)

type CustomField struct {
	Value         any   `json:"value"`
	CustomFieldId int32 `json:"custom_field_id"`

	MultiValue []CustomField `json:"-"`
}

func UnmarshalCustomFields(data []byte) (map[string]CustomField, error) {
	customFields := map[string]CustomField{}
	baseFields := map[string]any{}
	if err := json.Unmarshal(data, &baseFields); err != nil {
		return customFields, err
	}
	keys := maps.Keys(baseFields)
	for key := range keys {
		if !strings.HasPrefix(key, "custom_field_") {
			delete(baseFields, key)
		} else {
			data, err := json.Marshal(baseFields[key])
			if err != nil {
				return customFields, err
			}
			customField, err := unmarshalSingleCustomField(data)
			if err != nil {
				return customFields, err
			}
			customFields[key] = *customField
		}
	}
	return customFields, nil
}

func unmarshalSingleCustomField(data []byte) (*CustomField, error) {
	if len(data) == 0 {
		return nil, errors.New("custom_field is empty")
	}
	if data[0] == '[' {
		arrayCustomField := make([]CustomField, 0)
		if err := json.Unmarshal(data, &arrayCustomField); err != nil {
			return nil, err
		}
		return &CustomField{
			Value:         nil,
			MultiValue:    arrayCustomField,
			CustomFieldId: 0,
		}, nil
	}
	customField := CustomField{}
	if err := json.Unmarshal(data, &customField); err != nil {
		return nil, err
	}
	return &CustomField{
		Value:         customField.Value,
		CustomFieldId: customField.CustomFieldId,
	}, nil
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
	out := make(map[string]any)
	if err := json.Unmarshal(data, &out); err != nil {
		return nil, err
	}
	for _, k := range skippedFields {
		delete(out, k)
	}
	// Add extra fields back
	for k, v := range customFields {
		if v.Value != nil {
			out[k] = v
		}
		if v.MultiValue != nil {
			out[k] = v.MultiValue
		}
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
	return fmt.Appendf([]byte{}, "%d", p.Id), nil
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
	return err
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

func ErambaViewLink(base, tool string, id int32) string {
	filter := "sort%5Bcreated%5D=desc"
	return fmt.Sprintf("%s/%s/view/%s/%d?%s", base, tool, convertKebabToPascal(tool), id, filter)
}

func convertKebabToPascal(s string) string {
	parts := strings.Split(s, "-")
	for i := range parts {
		if len(parts[i]) > 0 {
			r := []rune(parts[i])
			r[0] = unicode.ToUpper(r[0])
			parts[i] = string(r)
		}
	}
	return strings.Join(parts, "")
}
