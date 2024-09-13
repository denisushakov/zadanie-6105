package models

import "time"

type TenderStatus string

const (
	TenderCreated   TenderStatus = "CREATED"
	TenderPublished TenderStatus = "PUBLISHED"
	TenderClosed    TenderStatus = "CLOSED"
)

type Tender struct {
	ID             string       `json:"id"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	ServiceType    string       `json:"service_type"`
	Status         TenderStatus `json:"status"`
	OrganizationID string       `json:"organization_id"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
}
