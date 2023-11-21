package service

import "strconv"

type DeleteServiceInput struct {
	ServiceId int
}

type DeleteServiceOutput struct {
}

func (c *Client) Delete(input *DeleteServiceInput) (*DeleteServiceOutput, error) {
	result := &DeleteServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Delete("services/" + strconv.Itoa(input.ServiceId))
	return result, err
}
