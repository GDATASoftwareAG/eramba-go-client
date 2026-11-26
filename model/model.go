package model

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type BusinessUnit struct {
	Id int32 `json:"id"`
}

func (b BusinessUnit) GetId() int32 {
	return b.Id
}

func (b BusinessUnit) Link(base string) string {
	return ErambaViewLink(base, "business-units", b.Id)
}

type UserOrGroup struct {
	ObjectKey string `json:"object_key"`
	Group     struct {
		Name string `json:"name"`
	} `json:"group,omitempty"`
	User struct {
		Name string `json:"name"`
	} `json:"user,omitempty"`
}

func (o UserOrGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.ObjectKey)
}

// UsersAndGroups is returned by the v2 API as {"users": [...], "groups": [...]}
// and must be sent as ["User-<id>", "Group-<id>"].
type UsersAndGroups struct {
	Users  []*User  `json:"users"`
	Groups []*Group `json:"groups"`
}

func (u UsersAndGroups) ObjectKeys() []string {
	keys := make([]string, 0, len(u.Users)+len(u.Groups))
	for _, user := range u.Users {
		keys = append(keys, fmt.Sprintf("User-%d", user.Id))
	}
	for _, group := range u.Groups {
		keys = append(keys, fmt.Sprintf("Group-%d", group.Id))
	}
	return keys
}

func (u UsersAndGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(u.ObjectKeys())
}

// UnmarshalJSON accepts the object form, an empty list/null (PHP serializes
// empty associative arrays as []) and the ["User-<id>", "Group-<id>"] form.
func (u *UsersAndGroups) UnmarshalJSON(data []byte) error {
	*u = UsersAndGroups{}
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) {
		return nil
	}
	if trimmed[0] == '[' {
		keys := []string{}
		if err := json.Unmarshal(trimmed, &keys); err != nil {
			return err
		}
		for _, key := range keys {
			kind, idStr, found := strings.Cut(key, "-")
			id, err := strconv.ParseInt(idStr, 10, 32)
			if !found || err != nil {
				return fmt.Errorf("invalid user/group reference %q", key)
			}
			switch kind {
			case "User":
				u.Users = append(u.Users, &User{Id: int32(id)})
			case "Group":
				u.Groups = append(u.Groups, &Group{Id: int32(id)})
			default:
				return fmt.Errorf("invalid user/group reference %q", key)
			}
		}
		return nil
	}
	type Alias UsersAndGroups // avoid recursion
	return json.Unmarshal(trimmed, (*Alias)(u))
}

type RiskAppetiteThreshold struct {
	Id    int32  `json:"risk_appetite_threshold_id"`
	Title string `json:"title,omitempty"`
}

type RiskThreat struct {
	Id              int32    `json:"id"`
	Title           string   `json:"name"`
	AssetMediaTypes []OnlyId `json:"asset_media_types"`
}

type Tag struct {
	Title string `json:"title"`
}

func (o Tag) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Title)
}

type RiskScore struct {
	Score int32 `json:"score"`
}

type RiskClassification struct {
	RiskClassificationId int32 `json:"risk_classification_id"`
}

func (o RiskClassification) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.RiskClassificationId)
}

func convertToTypeId(prefix, key string) (id string, typeId int32, err error) {
	key = strings.Replace(key, prefix, "", 1)

	splits := slices.DeleteFunc(strings.Split(key, "_"), func(s string) bool {
		return s == ""
	})
	if len(splits) < minRiskClassificationKeySplits {
		return "", 0, errors.New("not enough")
	}
	firstElement := splits[0]
	lastElement := splits[len(splits)-1]
	num, err := strconv.ParseInt(lastElement, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
		return firstElement, 0, err
	}

	return firstElement, int32(num), nil
}
