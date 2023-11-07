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

const DefaultTokenUrl = "https://launchboxhq.io/oauth/token"

const (
	LaunchboxTokenUrlEnvVar     = "LAUNCHBOX_TOKEN_URL"
	LaunchboxClientIdEnvVar     = "LAUNCHBOX_CLIENT_ID"
	LaunchboxClientSecretEnvVar = "LAUNCHBOX_CLIENT_SECRET"
	LaunchboxAccessToken        = "LAUNCHBOX_ACCESS_TOKEN"
)

// Credentials represent the access credentials
// provided to LaunchboxHQ
type Credentials struct {
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
