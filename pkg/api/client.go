package api

import (
	"context"
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/openapi"
)

type Client struct {
	Context context.Context
	Config  cfg.RecotemConfig
}

func NewClient(context context.Context, config cfg.RecotemConfig) Client {
	client := Client{}
	client.Context = context
	client.Config = config
	return client
}

func (c Client) newApiClient() (*openapi.ClientWithResponses, error) {
	provider, err := securityprovider.NewSecurityProviderApiKey("header", "Authorization", fmt.Sprintf("Token %s", c.Config.Token))
	if err != nil {
		return nil, err
	}
	client, err := openapi.NewClientWithResponses(c.Config.Url, openapi.WithRequestEditorFn(provider.Intercept))
	if err != nil {
		return nil, err
	}
	return client, nil
}
