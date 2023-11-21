package service

import (
	"github.com/launchboxio/launchbox-go-sdk/config"
	"time"
)

type Service struct {
	Id              int    `json:"id"`
	Name            string `json:"name,omitempty"`
	FullName        string `json:"full_name"`
	DefaultBranch   string `json:"default_branch,omitempty"`
	Visibility      string `json:"visibility,omitempty"`
	Language        string `json:"language,omitempty"`
	VcsConnectionId int    `json:"vcs_connection_id,omitempty"`

	DeploymentStrategy string `json:"deployment_strategy,omitempty"`
	UpdateStrategy     string `json:"update_strategy,omitempty"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Client struct {
	cnf *config.Config
}

func New(cnf *config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
