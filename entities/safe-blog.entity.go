package entities

import "time"

type SafeBlog struct {
	ID          string
	Slug        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	PublishedAt time.Time

	Title    string
	Summary  string
	Content  string
	CoverURL string

	AuthorID string
}
