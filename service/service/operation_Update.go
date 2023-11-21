package service

import "strconv"

type UpdateServiceInput struct {
	ServiceId int    `json:"-"`
	Name      string `json:"name"`
}

type UpdateServiceOutput struct {
	Service Service `json:"service"`
}

func (c *Client) Update(input *UpdateServiceInput) (*UpdateServiceOutput, error) {
	payload := struct {
		Service *UpdateServiceInput `json:"service"`
	}{
		Service: input,
	}
	result := &UpdateServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Patch("services/" + strconv.Itoa(input.ServiceId))
	return result, err
}
