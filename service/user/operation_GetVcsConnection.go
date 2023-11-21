package user

import "strconv"

type GetVcsConnectionInput struct {
	VcsConnectionId int
}

type GetVcsConnectionOutput struct {
	VcsConnection *VcsConnection `json:"vcs_connection"`
}

func (c *Client) Get(input *GetVcsConnectionInput) (*GetVcsConnectionOutput, error) {
	result := &GetVcsConnectionOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get("vcs_connections/" + strconv.Itoa(input.VcsConnectionId))
	return result, err
}
