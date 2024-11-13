package model

import "database/sql"

type Events struct {
	ID int `json:"id,omitempty"`
	Price float32 `json:"price,omitempty"`
	Date string `json:"date,omitempty"`
	DestinationID Destination `json:"destination,omitempty"`
	Traveler int `json:"traveler,omitempty"`
	Rating sql.NullFloat64 `json:"rating"`
}