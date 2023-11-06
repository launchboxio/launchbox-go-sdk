package project

type CreateProjectInput struct {
	Memory int
	Cpu    int
	Disk   int
	Name   string
}

type CreateProjectOutput struct {
	Project *Project `json:"project"`
}

func (c *Client) Create(input *CreateProjectInput) (*CreateProjectOutput, error) {
	result := &CreateProjectOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		Post("projects")
	return nil, nil
}
