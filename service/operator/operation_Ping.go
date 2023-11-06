package operator

type PingInput struct {
}

type PingOutput struct {
}

func (c *Client) Ping(input *PingInput) (*PingOutput, error) {
	return nil, nil
}
