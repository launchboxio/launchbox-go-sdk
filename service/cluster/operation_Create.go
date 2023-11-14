package cluster

type CreateClusterInput struct {
	Domain string `json:"domain"`
	Name   string `json:"name"`
}

type CreateClusterOutput struct {
	Cluster *Cluster `json:"cluster"`
}

func (c *Client) Create(input *CreateClusterInput) (*CreateClusterOutput, error) {
	payload := struct {
		Cluster *CreateClusterInput `json:"cluster"`
	}{
		Cluster: input,
	}
	result := &CreateClusterOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post("clusters")
	return result, err
}
