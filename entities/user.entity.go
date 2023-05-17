package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid; primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Provider    string `gorm:"type:varchar(48)"` // Third-party provider (e.g., Google, Facebook)
	ProviderID  string `gorm:"type:text"`        // ID provided by the third-party provider
	AccessToken string `gorm:"type:text"`        // Access token obtained from the third-party provider

	Username string `gorm:"type:varchar(100); unique"`
	Email    string `gorm:"type:varchar(100); unique"`
	Bio      string `gorm:"type:text; nullable"`
}
