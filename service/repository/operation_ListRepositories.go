package repository

type ListRepositoriesInput struct{}

func (*ListRepositoriesInput) ToQueryParams() map[string]string {
	return map[string]string{}
}

type ListRepositoriesOutput struct {
	Repositories []*Repository `json:"repositories"`
}

func (c *Client) GetRepositories(input *ListRepositoriesInput) (*ListRepositoriesOutput, error) {
	result := &ListRepositoriesOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetQueryParams(input.ToQueryParams()).
		SetResult(result).
		Get("repositories")
	return result, err
}
