package repository

type ListVcsConnectionsInput struct {
}

func (*ListVcsConnectionsInput) ToQueryParams() map[string]string {
	return map[string]string{}
}

type ListVcsConnectionsOutput struct {
	VcsConnections *VcsConnection `json:"vcs_connections"`
}

func (c *Client) ListVcsConnections(input *ListVcsConnectionsInput) (*ListVcsConnectionsOutput, error) {
	result := &ListVcsConnectionsOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetQueryParams(input.ToQueryParams()).
		SetResult(result).
		Patch("vcs_connections")
	return result, err
}
