package client

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleOAuth2 struct {
	Config *oauth2.Config
}

func NewGoogleOAuth2(baseURL, clientID, clientSecret string) *GoogleOAuth2 {
	cbURL := fmt.Sprintf("%s/auth/google/callback", baseURL)

	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  cbURL,
		Scopes:       []string{"email", "profile"},
		Endpoint:     google.Endpoint,
	}

	return &GoogleOAuth2{Config: config}
}
