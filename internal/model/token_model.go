package model

import "time"

type Token struct {
	ID        string    `json:"id"`
	Channel   string    `json:"channel"`
	Groups    string    `json:"groups"`
	Secret    string    `json:"secret"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Failure represents a failure response in the REST service.
type Failure struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
