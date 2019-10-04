package auth

import (
	"io/ioutil"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func ServiceAccount(credentialFile string, scope string) (*oauth2.Token, error) {
	b, err := ioutil.ReadFile(credentialFile)
	if err != nil {
		return nil, err
	}
	config, _ := google.JWTConfigFromJSON(b, scope)

	token, err := config.TokenSource(oauth2.NoContext).Token()
	if err != nil {
		return nil, err
	}

	return token, nil
}
