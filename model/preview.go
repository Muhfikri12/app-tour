package model

type Preview struct {
	ID int `json:"id,omitempty"`
	EventId Events `json:"event,omitempty"`
	Comment string `json:"comment,omitempty"`
}