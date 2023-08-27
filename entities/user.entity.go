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

	Fullname   string `gorm:"type:varchar(100); nullable;"`
	Username   string `gorm:"type:varchar(100); unique; not null"`
	Email      string `gorm:"type:varchar(100); unique; not null"`
	Bio        string `gorm:"type:text; nullable"`
	PictureURL string `gorm:"type:text; nullable"`

	// Social media fields
	WebsiteURL   string `gorm:"type:text; nullable"`
	GithubURL    string `gorm:"type:text; nullable"`
	LinkedinURL  string `gorm:"type:text; nullable"`
	InstagramURL string `gorm:"type:text; nullable"`
	TwitterURL   string `gorm:"type:text; nullable"`
	YoutubeURL   string `gorm:"type:text; nullable"`

	IsAdmin  bool `gorm:"type:bool; default:false"`
	IsTester bool `gorm:"type:bool; default:false"`

	Blogs []Blog `gorm:"foreignKey:AuthorID"` // has many relationship with blog
}
