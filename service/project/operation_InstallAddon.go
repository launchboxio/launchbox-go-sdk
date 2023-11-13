package project

import (
	"fmt"
	"strconv"
)

type InstallAddonInput struct {
	ProjectId int `json:"-"`
	AddonId   int `json:"addon_id"`
}

type InstallAddonOutput struct {
	ProjectAddon *ProjectAddon `json:"project_addon"`
}

func (c *Client) InstallAddon(input *InstallAddonInput) (*InstallAddonOutput, error) {
	payload := struct {
		ProjectAddon *InstallAddonInput `json:"project_addon"`
	}{
		ProjectAddon: input,
	}
	result := &InstallAddonOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post(fmt.Sprintf("projects/%s/addons", strconv.Itoa(input.ProjectId)))
	return result, nil
}
