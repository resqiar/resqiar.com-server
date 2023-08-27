package inputs

import "time"

type UpdateBlogInput struct {
	ID       string `validate:"required"`
	Title    string `validate:"omitempty,max=100"`
	Summary  string `validate:"omitempty,max=300"`
	Content  string `validate:"omitempty,max=50000"`
	CoverURL string `validate:"omitempty,url"`

	Prev string `validate:"omitempty,max=32"`
	Next string `validate:"omitempty,max=32"`
}

type SafeUpdateBlogInput struct {
	Title     string
	Summary   string
	Content   string
	CoverURL  string
	UpdatedAt time.Time

	Prev string `validate:"omitempty,max=32"`
	Next string `validate:"omitempty,max=32"`
}
