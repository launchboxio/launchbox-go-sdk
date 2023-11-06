package config

import "github.com/launchboxio/launchbox-go-sdk/launchbox"

type Config struct {
	Credentials *launchbox.Credentials
}

func Default() *Config {
	chain := launchbox.NewProviderChain()
	credentials := chain.Resolve()

	return &Config{
		Credentials: credentials,
	}
}
