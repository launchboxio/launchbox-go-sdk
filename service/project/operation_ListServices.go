package project

import (
	"fmt"
	"strconv"
)

type ListServicesInput struct {
	ProjectId int
}

type ListServicesOutput struct {
	ProjectServices []ProjectService `json:"service_subscriptions"`
}

func (c *Client) ListServices(input *ListServicesInput) (*ListServicesOutput, error) {
	result := &ListServicesOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get(fmt.Sprintf("projects/%s/services", strconv.Itoa(input.ProjectId)))
	return result, nil
}
