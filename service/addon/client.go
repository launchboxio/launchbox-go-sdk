package addon

import (
	"github.com/launchboxio/launchbox-go-sdk/config"
	"time"
)

type Addon struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	OciRegistry      string `json:"oci_registry"`
	OciVersion       string `json:"oci_version"`
	PullPolicy       string `json:"pull_policy"`
	ActivationPolicy string `json:"activation_policy"`

	// Admin only fields
	Status    string    `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`

	DefaultVersion *AddonVersion   `json:"default_version,omitempty"`
	Versions       []*AddonVersion `json:"versions,omitempty"`
}

type AddonVersion struct {
	Id        int    `json:"id"`
	Version   string `json:"version"`
	ClaimName string `json:"claim_name"`
	Group     string `json:"group"`
	Default   bool   `json:"default,omitempty"`
}

type Client struct {
	cnf *config.Config
}

func New(cnf *config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
