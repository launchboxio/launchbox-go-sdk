package service

import "strconv"

type CreateServiceInput struct {
	RepositoryId int `json:"repository_id"`

	Name               string `json:"name,omitempty"`
	DeploymentStrategy string `json:"deployment_strategy,omitempty"`
	UpdateStrategy     string `json:"update_strategy_strategy,omitempty"`
}

type CreateServiceOutput struct {
	Service *Service `json:"service"`
}

func (c *Client) Create(input *CreateServiceInput) (*CreateServiceOutput, error) {
	payload := struct {
		Service *CreateServiceInput `json:"service"`
	}{
		Service: input,
	}
	result := &CreateServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post("repositories/" + strconv.Itoa(input.RepositoryId) + "/services")
	return result, err
}
