package project

import (
	"github.com/go-resty/resty/v2"
	"github.com/launchboxio/launchbox-go-sdk/config"
	"time"
)

type Project struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Slug      string `json:"slug"`
	ClusterId int    `json:"cluster_id"`

	Memory int `json:"memory"`
	Cpu    int `json:"cpu"`
	Disk   int `json:"disk"`

	LastPausedAt  time.Time `json:"last_paused_at"`
	LastStartedAt time.Time `json:"last_started_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Host          string `json:"host"`
	CaCertificate string `json:"ca_certificate"`
}

type Client struct {
	cnf *config.Config

	httpClient *resty.Client
}

func New(cnf *config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
