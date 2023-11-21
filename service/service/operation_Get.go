package service

import "strconv"

type GetServiceInput struct {
	ServiceId int
}

type GetServiceOutput struct {
	Service *Service `json:"service"`
}

func (c *Client) Get(input *GetServiceInput) (*GetServiceOutput, error) {
	result := &GetServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Get("services/" + strconv.Itoa(input.ServiceId))
	return result, err
}
