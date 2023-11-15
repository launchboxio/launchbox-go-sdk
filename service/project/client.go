package project

import (
	"github.com/launchboxio/launchbox-go-sdk/config"
	"github.com/launchboxio/launchbox-go-sdk/service/addon"
	"time"
)

type Project struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Status    string `json:"status"`
	Slug      string `json:"slug"`
	ClusterId int    `json:"cluster_id"`

	Memory            int    `json:"memory"`
	Cpu               int    `json:"cpu"`
	Disk              int    `json:"disk"`
	KubernetesVersion string `json:"kubernetes_version,omitempty"`

	LastPausedAt  time.Time `json:"last_paused_at"`
	LastStartedAt time.Time `json:"last_started_at"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`

	Host          string `json:"host"`
	CaCertificate string `json:"ca_certificate"`
}

type ProjectAddon struct {
	Id      int         `json:"id"`
	AddonId int         `json:"addon_id"`
	Addon   addon.Addon `json:"addon,omitempty"`
}

type Client struct {
	cnf *config.Config
}

func New(cnf *config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
