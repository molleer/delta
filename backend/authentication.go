package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2"
)

var (
	gamma_url = os.Getenv("GAMMA_URL")
)

var client = oauth2.Config{
	ClientID:     os.Getenv("GAMMA_CLIENT_ID"),
	ClientSecret: os.Getenv("GAMMA_CLIENT_SECRET"),
	Endpoint: oauth2.Endpoint{
		AuthURL:   fmt.Sprintf("%s/api/oauth/authorize", os.Getenv("GAMMA_REDIRECT_URL")),
		TokenURL:  fmt.Sprintf("%s/api/oauth/token", gamma_url),
		AuthStyle: 0,
	},
	RedirectURL: os.Getenv("GAMMA_CALLBACK_URL"),
	Scopes:      nil,
}

func TokenIsValid(token string) bool{
	//TODO: Send request to Gamma and check token
	return token != ""
}

func GetLoginURL() string {
	return fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s",
			client.Endpoint.AuthURL,
			client.ClientID,
			client.RedirectURL)
}

func GetToken(grant string) (*oauth2.Token, error) {
	return client.Exchange(context.Background(), grant)
}
