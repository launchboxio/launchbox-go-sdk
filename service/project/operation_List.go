package project

type ListProjectInput struct {
}

type ListProjectOutput struct {
	Projects []Project `json:"projects"`
}

func (i *ListProjectInput) ToQueryParams() map[string]string {
	return map[string]string{}
}

func (c *Client) List(input *ListProjectInput) (*ListProjectOutput, error) {
	result := &ListProjectOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetQueryParams(input.ToQueryParams()).
		SetResult(result).
		Get("projects")
	return result, err
}
