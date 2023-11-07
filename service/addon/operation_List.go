package addon

type ListAddonInput struct{}

type ListAddonOutput struct {
	Addons []*Addon `json:"addons"`
}

func (i *ListAddonInput) ToQueryParams() map[string]string {
	return map[string]string{}
}

func (c *Client) List(input *ListAddonInput) (*ListAddonOutput, error) {
	result := &ListAddonOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetQueryParams(input.ToQueryParams()).
		SetResult(result).
		Get("addons")
	return result, err
}
