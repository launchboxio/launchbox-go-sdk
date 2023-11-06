package config

import "github.com/launchboxio/launchbox-go-sdk/launchbox"

type Config struct {
	Credentials *launchbox.Credentials

	Endpoint string
}

func Default() (*Config, error) {
	chain := launchbox.NewProviderChain()
	credentials := chain.Resolve()

	return &Config{
		Credentials: credentials,
		Endpoint:    "https://launchboxhq.io",
	}, nil
}

func DefaultWithEndpoint(endpoint string) (*Config, error) {
	cnf, err := Default()
	cnf.Endpoint = endpoint
	return cnf, err
}
