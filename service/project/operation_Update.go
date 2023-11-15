package project

import "strconv"

type UpdateProjectInput struct {
	ProjectId         int    `json:"-"`
	Memory            int    `json:"memory,omitempty"`
	Cpu               int    `json:"cpu,omitempty"`
	Disk              int    `json:"disk,omitempty"`
	Name              string `json:"name,omitempty"`
	Status            string `json:"status,omitempty"`
	CaCertificate     string `json:"ca_certificate,omitempty"`
	KubernetesVersion string `json:"kubernetes_version,omitempty"`
}

type UpdateProjectOutput struct {
	Project *Project `json:"project"`
}

func (c *Client) Update(input *UpdateProjectInput) (*UpdateProjectOutput, error) {
	payload := struct {
		Project *UpdateProjectInput `json:"project"`
	}{
		Project: input,
	}
	result := &UpdateProjectOutput{}
	client, err := c.cnf.GetClient()
	if err != nil {
		return nil, err
	}
	_, err = client.R().
		SetResult(result).
		SetBody(payload).
		Patch("projects/" + strconv.Itoa(input.ProjectId))
	return result, err
}
