package cluster

import "strconv"

type UpdateClusterInput struct {
	ClusterId       int    `json:"-"`
	AgentVersion    string `json:"agent_version,omitempty"`
	AgentIdentifier string `json:"agent_identifier,omitempty"`
	Version         string `json:"version,omitempty"`
	Provider        string `json:"provider,omitempty"`
	Region          string `json:"region,omitempty"`
	Status          string `json:"status,omitempty"`
	Domain          string `json:"domain,omitempty"`
}

type UpdateClusterOutput struct {
	Cluster Cluster `json:"cluster"`
}

func (c *Client) Update(input *UpdateClusterInput) (*UpdateClusterOutput, error) {
	payload := struct {
		Cluster *UpdateClusterInput `json:"cluster"`
	}{
		Cluster: input,
	}
	result := &UpdateClusterOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Patch("clusters/" + strconv.Itoa(input.ClusterId))
	return result, err
}
