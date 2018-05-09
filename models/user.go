package models

import "time"

// User - User database object
type User struct {
	ID           uint       `json:"id" gorm:"primary_key"`
	FirstName    string     `json:"firstName"`
	LastName     string     `json:"lastName"`
	EmailAddress string     `json:"emailAddress"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
}
