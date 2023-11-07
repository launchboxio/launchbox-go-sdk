package addon

import "strconv"

type GetAddonInput struct {
	AddonId int
}

type GetAddonOutput struct {
	Addon *Addon `json:"addon"`
}

func (c *Client) Get(input *GetAddonInput) (*GetAddonOutput, error) {
	result := &GetAddonOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get("addons/" + strconv.Itoa(input.AddonId))
	return result, err
}
