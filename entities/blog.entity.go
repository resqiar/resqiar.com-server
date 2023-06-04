package entities

import (
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type Blog struct {
	ID          string `gorm:"type:text; primaryKey; unique; not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time
	DeletedAt   gorm.DeletedAt

	Title     string `gorm:"type:varchar(100); not null"`
	Summary   string `gorm:"type:text"`
	Content   string `gorm:"type:text"`
	Published bool   `gorm:"type:bool; default:false"`
	CoverURL  string `gorm:"type:text"`

	AuthorID string `gorm:"type:text; not null"`
}

func (blog *Blog) BeforeCreate(tx *gorm.DB) error {
	var CUSTOM_ALPHABET = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	// Generate simple ID using nanoid,
	generatedID, err := gonanoid.Generate(CUSTOM_ALPHABET, 12)
	if err != nil {
		return err
	}

	// Then use that ID as a primary key.
	blog.ID = generatedID

	return nil
}
