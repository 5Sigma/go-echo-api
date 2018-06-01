package models

// Widget - model for the  widgets table
type Widget struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CreatorID   uint    `json:"-"`
	Creator     *User   `json:"creator"`
}
