package config

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/launchboxio/launchbox-go-sdk/launchbox"
)

type Config struct {
	Credentials *launchbox.Credentials

	Endpoint string

	client *resty.Client
}

func Default() (*Config, error) {
	chain := launchbox.NewProviderChain()
	credentials := chain.Resolve()

	return &Config{
		Credentials: credentials,
		Endpoint:    "https://launchboxhq.io",
	}, nil
}

func DefaultWithEndpoint(endpoint string) (*Config, error) {
	cnf, err := Default()
	cnf.Endpoint = endpoint
	return cnf, err
}

func (c *Config) GetClient() (*resty.Client, error) {
	client := resty.New()
	client.OnBeforeRequest(func(client *resty.Client, req *resty.Request) error {
		req.URL = fmt.Sprintf("%s/%s", c.Endpoint, req.URL)
		return nil
	})
	client.OnBeforeRequest(func(client *resty.Client, request *resty.Request) error {
		request.Header.Add("Authorization", "Bearer "+c.Credentials.AccessToken)
		return nil
	})
	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if response.IsError() {
			switch response.StatusCode() {
			case 401:
				return &launchbox.UnauthorizedError{}
			case 403:
				return &launchbox.ForbiddenError{}
			default:
				return &launchbox.GenericError{}
			}
		}
		return nil
	})
	client.OnError(func(request *resty.Request, err error) {
		//fmt.Println(err)
		//fmt.Println(request)
	})
	return client, nil
}
