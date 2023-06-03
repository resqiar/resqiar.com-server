package entities

import "time"

type SafeBlog struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time

	Title    string
	Summary  string
	Content  string
	CoverURL string

	AuthorID string
}
