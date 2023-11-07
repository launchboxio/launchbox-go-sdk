package launchbox

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessTokenCredentialsFetch(t *testing.T) {
	t.Setenv(LaunchboxAccessToken, "my-access-token")
	provider := NewAccessTokenCredentials()

	creds, err := provider.Fetch(context.TODO())
	assert.NoError(t, err)
	assert.Equal(t, "my-access-token", creds.AccessToken)
}
