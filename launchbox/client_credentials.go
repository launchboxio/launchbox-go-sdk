package launchbox

import (
	"context"
	"golang.org/x/oauth2/clientcredentials"
	"os"
)

type ClientCredentials struct {
	ClientId     string
	ClientSecret string
	TokenUrl     string
}

func NewClientCredentialsFromEnv() *ClientCredentials {
	return NewClientCredentials(
		os.Getenv(LaunchboxClientIdEnvVar),
		os.Getenv(LaunchboxClientSecretEnvVar),
	)
}

func NewClientCredentials(clientId string, clientSecret string) *ClientCredentials {
	creds := &ClientCredentials{
		ClientId:     clientId,
		ClientSecret: clientSecret,
		TokenUrl:     DefaultTokenUrl,
	}
	tokenUrl, hasTokenUrl := os.LookupEnv(LaunchboxTokenUrlEnvVar)
	if hasTokenUrl {
		creds.TokenUrl = tokenUrl
	}
	return creds
}

func (cc *ClientCredentials) Fetch(ctx context.Context) (Credentials, error) {
	clientCredentials := clientcredentials.Config{
		ClientID:     cc.ClientId,
		ClientSecret: cc.ClientSecret,
		TokenURL:     cc.TokenUrl,
	}
	token, err := clientCredentials.Token(ctx)
	if err != nil {
		return Credentials{}, err
	}
	return Credentials{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expires:      token.Expiry,
	}, nil
}
