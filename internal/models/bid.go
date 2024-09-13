package models

import "time"

type BidStatus string

const (
	BidCreated   BidStatus = "CREATED"
	BidPublished BidStatus = "PUBLISHED"
	BidCanceled  BidStatus = "CANCELED"
)

type Bid struct {
	ID        string    `json:"id"`
	TenderID  string    `json:"tender_id"`
	AuthorID  string    `json:"author_id"`
	Amount    float64   `json:"amount"`
	Status    BidStatus `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
