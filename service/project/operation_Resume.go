package project

import "strconv"

type ResumeProjectInput struct {
	ProjectId int
}

type ResumeProjectOutput struct {
	Project *Project `json:"project"`
}

func (c *Client) Resume(input *ResumeProjectInput) (*ResumeProjectOutput, error) {
	result := &ResumeProjectOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Post("projects/" + strconv.Itoa(input.ProjectId) + "/resume")
	return result, err
}
