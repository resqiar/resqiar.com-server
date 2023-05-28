package entities

import (
	"time"
)

type SafeUser struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	Username   string
	Email      string
	Bio        string
	PictureURL string
}
