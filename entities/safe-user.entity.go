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

	WebsiteURL   string
	GithubURL    string
	LinkedinURL  string
	InstagramURL string
	TwitterURL   string
	YoutubeURL   string

	IsTester bool
}
