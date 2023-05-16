package config

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleConfig() *oauth2.Config {
	googleConfig := &oauth2.Config{
		ClientID:     os.Getenv("G_CLIENT_ID"),
		ClientSecret: os.Getenv("G_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("G_REDIRECT"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	return googleConfig
}
