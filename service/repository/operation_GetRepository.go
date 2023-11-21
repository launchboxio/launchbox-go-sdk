package repository

import "strconv"

type GetRepositoryInput struct {
	RepositoryId int
}

type GetRepositoryOutput struct {
	Repository *Repository `json:"repository"`
}

func (c *Client) GetRepository(input *GetRepositoryInput) (*GetRepositoryOutput, error) {
	result := &GetRepositoryOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get("repositories/" + strconv.Itoa(input.RepositoryId))
	return result, err
}
