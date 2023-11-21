package service

type CreateServiceInput struct {
	Name            string `json:"name,omitempty"`
	FullName        string `json:"full_name"`
	VcsConnectionId int    `json:"vcs_connection_id"`
}

type CreateServiceOutput struct {
	Service *Service `json:"service"`
}

func (c *Client) Create(input *CreateServiceInput) (*CreateServiceOutput, error) {
	payload := struct {
		Service *CreateServiceInput `json:"service"`
	}{
		Service: input,
	}
	result := &CreateServiceOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post("services")
	return result, err
}
