package entities

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string `gorm:"type:uuid; primaryKey; default:gen_random_uuid()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Provider   string `gorm:"type:varchar(48); not null"` // Third-party provider (e.g., Google, Facebook)
	ProviderID string `gorm:"type:text"`                  // ID provided by the third-party provider

	Username string `gorm:"type:varchar(100); unique; not null"`
	Email    string `gorm:"type:varchar(100); unique; not null"`
	Bio      string `gorm:"type:text; nullable"`

	IsAdmin bool `gorm:"type:bool; default:false"`
}
