package project

import (
	"strconv"
)

type GetProjectInput struct {
	ProjectId int
}

type GetProjectOutput struct {
	Project *Project `json:"project"`
}

func (c *Client) Get(input *GetProjectInput) (*GetProjectOutput, error) {
	result := &GetProjectOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get("projects/" + strconv.Itoa(input.ProjectId))
	return result, err
}
