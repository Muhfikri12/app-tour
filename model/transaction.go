package model

import "time"

type Transaction struct {
	ID            int    `json:"id,omitempty"`
	EventID       int `json:"event_id,omitempty"`
	Name          string `json:"name,omitempty" validate:"required,alphanum"`
	Phone         string `json:"phone,omitempty" validate:"required"`
	Email         string `json:"email,omitempty" validate:"required,email"`
	EmailConfirm  string `json:"email_confirm,omitempty" validate:"required,eqfield=Email"`
	Message       string `json:"message,omitempty" validate:"required,alphanum,min=30,max=255"`
	CreatedAt    time.Time `json:"created_at,omitempty" `
}