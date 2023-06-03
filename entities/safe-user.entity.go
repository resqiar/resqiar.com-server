package entities

import (
	"time"
)

type SafeUser struct {
	ID        string
	CreatedAt time.Time

	Username   string
	Bio        string
	PictureURL string
}
