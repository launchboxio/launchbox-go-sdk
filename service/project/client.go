package project

import "github.com/launchboxio/launchbox-go-sdk/config"

type Project struct {
}

type Client struct {
	cnf config.Config
}

func New(cnf config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
