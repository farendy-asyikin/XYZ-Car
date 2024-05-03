package models

import "time"

type Stock struct {
	ID        uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Brand     string    `gorm:"type:varchar(255);not null" json:"brand"`
	Type      string    `gorm:"type:varchar(255);not null" json:"type"`
	Price     int       `gorm:"type:int;not null" json:"price"`
	Amount    int       `gorm:"type:int;not null" json:"amount"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
