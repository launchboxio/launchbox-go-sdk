package project

import (
	"github.com/launchboxio/launchbox-go-sdk/service/addon"
	"strconv"
)

type GetProjectManifestInput struct {
	ProjectId int
}

type GetProjectManifestOutput struct {
	Manifest *Manifest `json:"project"`
}

type Manifest struct {
	Id                 int                 `json:"id"`
	Name               string              `json:"name"`
	Status             string              `json:"status,omitempty"`
	Slug               string              `json:"slug"`
	ClusterId          int                 `json:"cluster_id"`
	Memory             int                 `json:"memory"`
	Cpu                int                 `json:"cpu"`
	Disk               int                 `json:"disk"`
	KubernetesVersion  string              `json:"kubernetes_version,omitempty"`
	AddonSubscriptions []AddonSubscription `json:"addon_subscriptions"`
	User               User                `json:"user"`
}

type User struct {
	Email string `json:"email"`
}

type AddonSubscription struct {
	Id      int          `json:"id"`
	AddonId int          `json:"addon_id"`
	Name    string       `json:"name,omitempty"`
	Status  string       `json:"status,omitempty"`
	Addon   *addon.Addon `json:"addon"`
}

func (c *Client) GetManifest(input *GetProjectManifestInput) (*GetProjectManifestOutput, error) {
	result := &GetProjectManifestOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get("projects/" + strconv.Itoa(input.ProjectId) + "/manifest")
	return result, err
}
