package cluster

import "strconv"

type GetClusterInput struct {
	ClusterId int
}

type GetClusterOutput struct {
	Cluster *Cluster `json:"cluster"`
}

func (c *Client) Get(input *GetClusterInput) (*GetClusterOutput, error) {
	result := &GetClusterOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get("clusters/" + strconv.Itoa(input.ClusterId))
	return result, err
}
