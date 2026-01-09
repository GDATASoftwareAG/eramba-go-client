package model

import (
	"encoding/json"
	"strconv"
)

const (
	UseAttachmentUrl UseAttachment = 2
)

type UseAttachment int32

func (t UseAttachment) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Itoa(int(t))), nil
}

type SecurityPolices []*SecurityPolicy

func (p SecurityPolices) MarshalJSON() ([]byte, error) {
	list := extractPatchListId(p)
	return json.Marshal(list)
}

type SecurityPolicy struct {
	Id                           int32                  `json:"id"`
	Title                        string                 `json:"index"`
	Description                  string                 `json:"short_description"`
	SecurityPolicyDocumentTypeID int32                  `json:"security_policy_document_type_id"`
	Projects                     Projects               `json:"projects"`
	SecurityServices             []SecurityService      `json:"security_services"`
	RelatedDocuments             SecurityPolices        `json:"related_documents"`
	Tags                         []Tag                  `json:"tags"`
	PolicyReviewerContacts       []UserOrGroup          `json:"collaborators"`
	GrcContacts                  []UserOrGroup          `json:"owners"`
	Permission                   string                 `json:"permission"`
	Status                       int32                  `json:"status"`
	AssetLabelID                 *AssetLabelId          `json:"asset_label_id"`
	Version                      string                 `json:"version"`
	PublishedDate                *ErambaDate            `json:"published_date"`
	NextReviewDate               *ErambaDate            `json:"next_review_date"`
	UseAttachments               UseAttachment          `json:"use_attachments"`
	Url                          string                 `json:"url"`
	DocumentDescription          string                 `json:"description"`
	Reviews                      []SecurityPolicyReview `json:"security_policy_reviews"`
}

func (p *SecurityPolicy) GetId() int32 {
	return p.Id
}

func (p *SecurityPolicy) Link(base string) string {
	return buildLink(base, "security-policies", p.Id)
}

var SecurityPolicySkippedFields = []string{
	"id",
	"security_policy_reviews",
}

func (p *SecurityPolicy) MarshalJSON() ([]byte, error) {
	type Alias SecurityPolicy
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, SecurityPolicySkippedFields)
}

type SecurityPolicyReview struct {
	Id                int32         `json:"id"`
	ForeignKey        int32         `json:"foreign_key"`
	PlannedDate       *ErambaDate   `json:"planned_date"`
	ActualDate        *ErambaDate   `json:"actual_date"`
	Completed         int32         `json:"completed"`
	Version           string        `json:"version"`
	UseAttachments    UseAttachment `json:"use_attachments"`
	Url               string        `json:"url"`
	Description       string        `json:"description"`
	PolicyDescription string        `json:"policy_description"`
	Reviewers         []UserOrGroup `json:"reviewers"`
}

var SecurityPolicyReviewSkippedFields = []string{
	"id",
}

func (p *SecurityPolicyReview) MarshalJSON() ([]byte, error) {
	type Alias SecurityPolicyReview
	aux := Alias(*p)
	return MarshalWithSkippingFields(aux, map[string]CustomField{}, SecurityPolicyReviewSkippedFields)
}

type PolicyDescription struct {
	ContentType string `json:"content_type"`
}
