package model

type Events struct {
	ID int `json:"id,omitempty"`
	Price float32 `json:"price,omitempty"`
	Date string `json:"date,omitempty"`
	Rating *float64 `json:"rating"`
	DestinationID Destination `json:"destination,omitempty"`
	
}