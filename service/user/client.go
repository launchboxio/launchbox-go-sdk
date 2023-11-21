package user

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

type Client struct {
	cnf *config.Config
}

func New(cnf *config.Config) *Client {
	return &Client{
		cnf: cnf,
	}
}
