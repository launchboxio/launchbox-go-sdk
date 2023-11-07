package config

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/launchboxio/launchbox-go-sdk/launchbox"
	"os"
)

type Config struct {
	CredentialsProvider launchbox.CredentialsProvider

	Endpoint string

	client *resty.Client

	credentials *launchbox.Credentials
}

func Default() (*Config, error) {
	chain := launchbox.NewProviderChain()
	provider := chain.Resolve()

	endpoint, hasEndpoint := os.LookupEnv("LAUNCHBOX_URL")
	if hasEndpoint {
		return &Config{
			CredentialsProvider: provider,
			Endpoint:            endpoint,
		}, nil
	}
	return &Config{
		CredentialsProvider: provider,
		Endpoint:            "https://launchboxhq.io/api/v1",
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
		if c.credentials == nil || c.credentials.Expired() {
			creds, err := c.CredentialsProvider.Fetch(context.TODO())
			if err != nil {
				return err
			}
			c.credentials = &creds
		}

		request.Header.Add("Authorization", "Bearer "+c.credentials.AccessToken)
		return nil
	})
	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		if response.IsError() {
			switch response.StatusCode() {
			case 401:
				return &launchbox.UnauthorizedError{}
			case 403:
				return &launchbox.ForbiddenError{}
			case 404:
				return &launchbox.ResourceNotFoundError{}
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
