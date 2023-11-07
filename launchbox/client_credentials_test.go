package launchbox

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestClientCredentialsFromEnv(t *testing.T) {
	t.Setenv(LaunchboxClientIdEnvVar, "client-id")
	t.Setenv(LaunchboxClientSecretEnvVar, "client-secret")

	creds := NewClientCredentialsFromEnv()
	assert.Equal(t, "client-id", creds.ClientId)
	assert.Equal(t, "client-secret", creds.ClientSecret)
}

func TestClientCredentialsSetsTokenUrl(t *testing.T) {
	t.Setenv(LaunchboxTokenUrlEnvVar, "https://custom.io/oauth/token")
	creds := NewClientCredentialsFromEnv()
	assert.Equal(t, "https://custom.io/oauth/token", creds.TokenUrl)
}
