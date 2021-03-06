package models

import "time"

// User - User database object
type User struct {
	id           uint
	FirstName    string    `json:"firstName"`
	LastName     string    `json:"lastName"`
	EmailAddress string    `json:"emailAddress"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
