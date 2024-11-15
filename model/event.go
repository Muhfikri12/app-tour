package model

type Events struct {
	ID int `json:"id,omitempty"`
	Price float32 `json:"price,omitempty"`
	Date string `json:"date,omitempty"`
	Rating *float64 `json:"rating,omitempty"`
	DestinationID *Destination `json:"destination,omitempty"`	
}

type EventPlan struct {
	ID int `json:"id,omitempty"`
	Event *Events `json:"event,omitempty"`
	FirstPlan string `json:"first_plan,omitempty"`
	SecondPlan string `json:"second_plan,omitempty"`
	ThirdPlan string `json:"third_plan,omitempty"`
	FourthPlan string `json:"fourth_plan,omitempty"`
	FifthPlan string `json:"fifth_plan,omitempty"`
}