package models

import "time"

// APIKey - User API access tokens
type APIKey struct {
	ID        uint      `json:"id"`
	Key       string    `json:"key"`
	CreatedAt time.Time `json:"createdAt"`
	UserID    uint      `json:"userId"`
	User      User      `json:"user"`
}
