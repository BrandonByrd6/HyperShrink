package models

import "time"

type Url struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	LongUrl   string    `gorm:"not null" json:"long_url"`
	ShortUrl  string    `gorm:"unique;not null" json:"short_url"`
	UserId    string    `gorm:"not null" json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
