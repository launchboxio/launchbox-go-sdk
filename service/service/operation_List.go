package service

type ListServiceInput struct {
}

type ListServiceOutput struct {
	Services []*Service `json:"services"`
}

func (i *ListServiceInput) ToQueryParams() map[string]string {
	return map[string]string{}
}

func (c *Client) List(input *ListServiceInput) (*ListServiceOutput, error) {
	result := &ListServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetQueryParams(input.ToQueryParams()).
		SetResult(result).
		Get("services")
	return result, err
}
