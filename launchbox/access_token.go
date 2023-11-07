package launchbox

import (
	"context"
	"os"
)

type AccessTokenCredentials struct {
}

func NewAccessTokenCredentials() *AccessTokenCredentials {
	return &AccessTokenCredentials{}
}

func (atc *AccessTokenCredentials) Fetch(ctx context.Context) (Credentials, error) {
	return Credentials{
		AccessToken: os.Getenv(LaunchboxAccessToken),
	}, nil
}
