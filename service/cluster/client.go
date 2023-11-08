package cluster

import (
	"github.com/launchboxio/launchbox-go-sdk/config"
	"time"
)

type Cluster struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Region   string `json:"region"`
	Version  string `json:"version"`
	Provider string `json:"provider"`
	Status   string `json:"status"`

	// Fields that are only returned for admins
	CreatedAt       time.Time `json:"created_at,omitempty"`
	UpdatedAt       time.Time `json:"updated_at,omitempty"`
	AgentConnected  bool      `json:"agent_connected,omitempty"`
	AgentLastPing   time.Time `json:"agent_last_ping,omitempty"`
	AgentIdentifier string    `json:"agent_identifier,omitempty"`
	AgentVersion    string    `json:"agent_version,omitempty"`

	Credentials *ClientCredentials `json:"oauth_application,omitempty"`
}

type ClientCredentials struct {
	ClientId     string `json:"uid"`
	ClientSecret string `json:"secret"`
}

type Client struct {
	cnf *config.Config
}

func New(cnf *config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
