package models

import "time"

type Review struct {
	ID        string    `json:"id"`
	BidID     string    `json:"bid_id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}