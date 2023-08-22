package inputs

type UpdateUserInput struct {
	ID           string `validate:"required"`
	Username     string `validate:"omitempty,max=100"`
	Bio          string `validate:"omitempty,max=300"`
	PictureURL   string `validate:"omitempty,url"`
	WebsiteURL   string `validate:"omitempty,url"`
	GitHubURL    string `validate:"omitempty,url"`
	LinkedInURL  string `validate:"omitempty,url"`
	InstagramURL string `validate:"omitempty,url"`
	TwitterURL   string `validate:"omitempty,url"`
	YoutubeURL   string `validate:"omitempty,url"`
}
