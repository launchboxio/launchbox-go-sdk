package cluster

type ListClusterInput struct {
}

type ListClusterOutput struct {
	Clusters []*Cluster `json:"clusters"`
}

func (i *ListClusterInput) ToQueryParams() map[string]string {
	return map[string]string{}
}

func (c *Client) List(input *ListClusterInput) (*ListClusterOutput, error) {
	result := &ListClusterOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetQueryParams(input.ToQueryParams()).
		SetResult(result).
		Get("clusters")
	return result, err
}
