package addon

type CreateAddonInput struct {
	Name             string `json:"name"`
	OciRegistry      string `json:"oci_registry"`
	OciVersion       string `json:"oci_version"`
	PullPolicy       string `json:"pull_policy"`
	ActivationPolicy string `json:"activation_policy"`
}

type CreateAddonOutput struct {
	Addon *Addon `json:"addon"`
}

func (c *Client) Create(input *CreateAddonInput) (*CreateAddonOutput, error) {
	payload := struct {
		Addon *CreateAddonInput `json:"addon"`
	}{
		Addon: input,
	}
	result := &CreateAddonOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post("addons")
	return result, err
}
