package model

type Response struct {
	StatusCode int `json:"status_code,omitempty"`
	Status string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Page int `json:"page,omitempty"`
	Limit int `json:"limit,omitempty"`
	TotalData int `json:"total_data,omitempty"`
	TotalPage int `json:"total_page,omitempty"`
	Data interface{} `json:"data,omitempty"`
}