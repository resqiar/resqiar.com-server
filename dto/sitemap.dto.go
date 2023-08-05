package dto

import "time"

type SitemapOutput struct {
	AuthorUsername string
	Slug           string
	UpdatedAt      time.Time
}
