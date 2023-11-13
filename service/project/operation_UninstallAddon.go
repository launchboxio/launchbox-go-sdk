package project

import (
	"fmt"
	"strconv"
)

type UninstallAddonInput struct {
	ProjectId      int
	SubscriptionId int
}

type UninstallAddonOutput struct {
}

func (c *Client) UninstallAddon(input *UninstallAddonInput) (*UninstallAddonOutput, error) {
	result := &UninstallAddonOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Delete(fmt.Sprintf("projects/%s/addons/%s", strconv.Itoa(input.ProjectId), strconv.Itoa(input.SubscriptionId)))
	return nil, nil
}
