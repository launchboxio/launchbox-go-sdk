package launchbox

import "os"

type ProviderChain struct {
}

func NewProviderChain() *ProviderChain {
	return &ProviderChain{}
}

func (pc *ProviderChain) Resolve() CredentialsProvider {
	_, hasClientSecret := os.LookupEnv("LAUNCHBOX_CLIENT_SECRET")
	if hasClientSecret {
		return NewClientCredentialsFromEnv()
	}
	return NewAccessTokenCredentials()
}
