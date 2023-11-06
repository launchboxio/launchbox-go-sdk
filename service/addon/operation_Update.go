package addon

import "strconv"

type UpdateAddonInput struct {
	AddonId          int    `json:"-"`
	OciRegistry      string `json:"oci_registry,omitempty"`
	OciVersion       string `json:"oci_version,omitempty"`
	ActivationPolicy string `json:"activation_policy,omitempty"`
	PullPolicy       string `json:"pull_policy,omitempty"`
}

type UpdateAddonOutput struct {
	Addon *Addon `json:"addon"`
}

func (c *Client) Update(input *UpdateAddonInput) (*UpdateAddonOutput, error) {
	payload := struct {
		Addon *UpdateAddonInput `json:"addon"`
	}{
		Addon: input,
	}
	result := &UpdateAddonOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Patch("addons/" + strconv.Itoa(input.AddonId))
	return result, err
}
