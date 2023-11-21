package repository

type CreateRepositoryInput struct {
	VcsConnectionId int    `json:"vcs_connection_id,omitempty"`
	Repository      string `json:"repository"`
}

type CreateRepositoryOutput struct {
	Repository *Repository `json:"repository"`
}

func (c *Client) CreateRepository(input *CreateRepositoryInput) (*CreateRepositoryOutput, error) {
	payload := struct {
		Repository *CreateRepositoryInput `json:"repository"`
	}{
		Repository: input,
	}
	result := &CreateRepositoryOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post("repositories")
	return result, err
}
