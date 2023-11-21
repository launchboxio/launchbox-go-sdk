package repository

import "strconv"

type DeleteRepositoryInput struct {
	RepositoryId int `json:"-"`
}

type DeleteRepositoryOutput struct {
}

func (c *Client) DeleteRepository(input *DeleteRepositoryInput) (*DeleteRepositoryOutput, error) {
	result := &DeleteRepositoryOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Delete("repositories/" + strconv.Itoa(input.RepositoryId))
	return result, err
}
