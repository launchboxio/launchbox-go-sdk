package project

import "strconv"

type PauseProjectInput struct {
	ProjectId int
}

type PauseProjectOutput struct {
	Project *Project `json:"project"`
}

func (c *Client) Pause(input *PauseProjectInput) (*PauseProjectOutput, error) {
	result := &PauseProjectOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Post("projects/" + strconv.Itoa(input.ProjectId) + "/pause")
	return result, err
}
