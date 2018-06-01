package models

import "time"

// APIKey - User API access tokens
type APIKey struct {
	ID         uint      `json:"-"`
	Key        string    `json:"key"`
	CreatedAt  time.Time `json:"createdAt"`
	UserID     uint      `json:"-"`
	User       User      `json:"-"`
	SessionKey bool      `json:"-"`
	Expiration time.Time `json:"expirationDate"`
}
