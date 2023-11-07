package launchbox

import (
	"context"
	"time"
)

type CredentialsType string

const (
	ClientCredentialsType CredentialsType = "CLIENT_CREDENTIALS"
	AccessTokenType                       = "ACCESS_TOKEN"
)

type Credentials struct {
	Type         CredentialsType
	ClientId     string
	ClientSecret string
	IssuerUrl    string
	AccessToken  string
	RefreshToken string
	Expires      time.Time
}

type CredentialsProvider interface {
	Fetch(ctx context.Context) (Credentials, error)
}

// Expired checks if a credential set has expired
func (c Credentials) Expired() bool {
	return !c.Expires.After(time.Now())
}
