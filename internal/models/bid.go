package models

import "time"

type Bid struct {
	ID        string    `json:"id"`
	TenderID  string    `json:"tender_id"`
	AuthorID  string    `json:"author_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
