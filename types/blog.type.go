package types

type GetBlogOpts struct {
	UseID          string
	BlogAuthor     string
	BlogSlug       string
	IncludeContent bool
	Published      bool
}

type BlogDetailOpts struct {
	*GetBlogOpts
	ReturnHTML bool
}
