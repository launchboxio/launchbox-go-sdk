package repository

import (
	"github.com/launchboxio/launchbox-go-sdk/config"
	"time"
)

type VcsConnection struct {
	Id       int       `json:"id"`
	Method   string    `json:"method,omitempty"`
	Provider string    `json:"provider"`
	Host     string    `json:"host,omitempty"`
	Email    string    `json:"email,omitempty"`
	Uid      string    `json:"uid,omitempty"`
	Expiry   time.Time `json:"expiry,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Repository struct {
	Id            int    `json:"id"`
	RepositoryUrl string `json:"repository_url,omitempty"`
	Repository    string `json:"repository"`

	DefaultBranch string `json:"default_branch,omitempty"`
	Visibility    string `json:"visibility,omitempty"`
	Language      string `json:"language,omitempty"`

	DefaultDeploymentStrategy string `json:"default_deployment_strategy,omitempty"`
	DefaultUpdateStrategy     string `json:"default_update_strategy,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	VcsConnectionId int `json:"vcs_connection_id,omitempty"`
}

type Client struct {
	cnf *config.Config
}

func New(cnf *config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
