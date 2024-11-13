package model

type Transaction struct {
	ID int `json:"id,omitempty"`
	EventId Events `json:"event,omitempty"`
}