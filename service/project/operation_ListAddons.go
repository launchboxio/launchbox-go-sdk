package project

import (
	"fmt"
	"strconv"
)

type ListAddonsInput struct {
	ProjectId int
}

type ListAddonsOutput struct {
	ProjectAddons []ProjectAddon `json:"project_addons"`
}

func (c *Client) ListAddons(input *ListAddonsInput) (*ListAddonsOutput, error) {
	result := &ListAddonsOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get(fmt.Sprintf("projects/%s/addons", strconv.Itoa(input.ProjectId)))
	return result, nil
}
