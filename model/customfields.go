package model

import (
	"encoding/json"
	"errors"
	"maps"
	"strings"
)

type CustomFields map[string]CustomField

type CustomField struct {
	Value         any   `json:"value"`
	CustomFieldId int32 `json:"custom_field_id"`

	MultiValue []CustomField `json:"-"`
}

func (c CustomFields) SetString(key, value string) {
	if field, exists := c[key]; !exists {
		c[key] = CustomField{
			Value: value,
		}
	} else {
		field.Value = value
	}
}

func (c CustomFields) GetString(key string) (string, bool) {
	if value, exists := c[key]; exists {
		str, isStr := value.Value.(string)
		if isStr {
			return str, true
		}
	}
	return "", false
}

func (c CustomFields) SetInt(key string, value int) {
	if field, exists := c[key]; !exists {
		c[key] = CustomField{
			Value: value,
		}
	} else {
		field.Value = value
	}
}

func (c CustomFields) GetInt(key string) (int, bool) {
	if value, exists := c[key]; exists {
		integer, isInt := value.Value.(int)
		if isInt {
			return integer, true
		}
	}
	return 0, false
}

func UnmarshalCustomFields(data []byte) (CustomFields, error) {
	customFields := CustomFields{}
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
	customFields CustomFields,
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
