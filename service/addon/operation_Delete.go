package addon

import "strconv"

type DeleteAddonInput struct {
	AddonId int
}

type DeleteAddonOutput struct {
}

func (c *Client) Delete(input *DeleteAddonInput) (*DeleteAddonOutput, error) {
	result := &DeleteAddonOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Delete("addons/" + strconv.Itoa(input.AddonId))
	return result, err
}
