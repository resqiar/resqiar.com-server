package inputs

type CreateBlogInput struct {
	Title    string `validate:"required,max=100"`
	Summary  string `validate:"max=300"`
	Content  string `validate:"max=50000"`
	CoverURL string `validate:"omitempty,url"`
}
