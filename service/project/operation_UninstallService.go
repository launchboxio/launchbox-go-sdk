package project

import (
	"fmt"
	"strconv"
)

type UninstallServiceInput struct {
	ProjectId      int
	SubscriptionId int
}

type UninstallServiceOutput struct {
}

func (c *Client) UninstallService(input *UninstallServiceInput) (*UninstallServiceOutput, error) {
	result := &UninstallServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Delete(fmt.Sprintf("projects/%s/services/%s", strconv.Itoa(input.ProjectId), strconv.Itoa(input.SubscriptionId)))
	return nil, nil
}
