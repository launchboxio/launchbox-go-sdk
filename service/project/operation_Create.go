package project

type CreateProjectInput struct {
	Memory int    `json:"memory,omitempty"`
	Cpu    int    `json:"cpu,omitempty"`
	Disk   int    `json:"disk,omitempty"`
	Name   string `json:"name"`
}

type CreateProjectOutput struct {
	Project *Project `json:"project"`
}

func (c *Client) Create(input *CreateProjectInput) (*CreateProjectOutput, error) {
	payload := struct {
		Project *CreateProjectInput `json:"project"`
	}{
		Project: input,
	}
	result := &CreateProjectOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Post("projects")
	return result, err
}
