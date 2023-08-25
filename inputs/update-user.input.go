package inputs

type UpdateUserInput struct {
	Fullname     string `validate:"omitempty,max=100"`
	Username     string `validate:"omitempty,min=3,max=100,username"`
	Bio          string `validate:"max=300"`
	PictureURL   string `validate:"omitempty,media_url"`
	WebsiteURL   string `validate:"omitempty,media_url"`
	GithubURL    string `validate:"omitempty,media_url"`
	LinkedinURL  string `validate:"omitempty,media_url"`
	InstagramURL string `validate:"omitempty,media_url"`
	TwitterURL   string `validate:"omitempty,media_url"`
	YoutubeURL   string `validate:"omitempty,media_url"`
}
