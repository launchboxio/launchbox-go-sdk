package cluster

import "strconv"

type DeleteClusterInput struct {
	ClusterId int
}

type DeleteClusterOutput struct {
}

func (c *Client) Delete(input *DeleteClusterInput) (*DeleteClusterOutput, error) {
	result := &DeleteClusterOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Delete("clusters/" + strconv.Itoa(input.ClusterId))
	return result, err
}
