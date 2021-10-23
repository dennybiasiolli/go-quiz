/**
 * inspired from https://github.com/meehow/go-django-hashers
 */
package auth

import (
	"github.com/dennybiasiolli/go-quiz/common"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func getGoogleOauth2Config() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  common.GOOGLE_OAUTH2_DEFAULT_REDIRECT_URL,
		ClientID:     common.GOOGLE_OAUTH2_CLIENT_ID,
		ClientSecret: common.GOOGLE_OAUTH2_CLIENT_SECRET,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}
