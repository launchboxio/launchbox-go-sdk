package project

import (
	"fmt"
	"strconv"
)

type InstallServiceInput struct {
	ProjectId int `json:"-"`
	ServiceId int `json:"service_id"`
}

type InstallServiceOutput struct {
	ProjectService *ProjectService `json:"service_subscription"`
}

func (c *Client) InstallService(input *InstallServiceInput) (*InstallServiceOutput, error) {
	payload := struct {
		ProjectService *InstallServiceInput `json:"service_subscription"`
	}{
		ProjectService: input,
	}
	result := &InstallServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post(fmt.Sprintf("projects/%s/services", strconv.Itoa(input.ProjectId)))
	return result, nil
}
