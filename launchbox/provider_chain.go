package launchbox

import "os"

type ProviderChain struct {
}

func NewProviderChain() *ProviderChain {
	return &ProviderChain{}
}

func (pc *ProviderChain) Resolve() *Credentials {
	clientSecret, hasClientSecret := os.LookupEnv("LAUNCHBOX_CLIENT_SECRET")
	clientId, _ := os.LookupEnv("LAUNCHBOX_CLIENT_ID")
	accessToken, _ := os.LookupEnv("LAUNCHBOX_ACCESS_TOKEN")
	refreshToken, _ := os.LookupEnv("LAUNCHBOX_REFRESH_TOKEN")

	// For now, we just support Client Credentials, or Access Token
	// TODO: Support sourcing from file as well
	if hasClientSecret {
		return &Credentials{
			ClientId:     clientId,
			ClientSecret: clientSecret,
			Type:         ClientCredentialsType,
		}
	}
	return &Credentials{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Type:         AccessTokenType,
	}
}
