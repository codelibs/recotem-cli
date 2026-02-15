package api

import (
	"context"
	"fmt"
	"net/http"

	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/openapi"
)

type Client struct {
	Context context.Context
	Config  cfg.RecotemConfig
}

func NewClient(ctx context.Context, config cfg.RecotemConfig) Client {
	return Client{
		Context: ctx,
		Config:  config,
	}
}

// newApiClient creates an authenticated OpenAPI client.
// Authentication priority: API key flag > JWT access token > legacy Token
func (c Client) newApiClient() (*openapi.ClientWithResponses, error) {
	authFn := func(ctx context.Context, req *http.Request) error {
		// Priority 1: API key
		if c.Config.ApiKey != "" {
			req.Header.Set("X-API-Key", c.Config.ApiKey)
			return nil
		}
		// Priority 2: JWT access token
		if c.Config.AccessToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Config.AccessToken))
			return nil
		}
		// Priority 3: Legacy token
		if c.Config.Token != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Token %s", c.Config.Token))
			return nil
		}
		return nil
	}

	client, err := openapi.NewClientWithResponses(
		c.Config.Url,
		openapi.WithRequestEditorFn(authFn),
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// newUnauthenticatedClient creates an OpenAPI client without authentication (for login, ping, etc.)
func (c Client) newUnauthenticatedClient() (*openapi.ClientWithResponses, error) {
	client, err := openapi.NewClientWithResponses(c.Config.Url)
	if err != nil {
		return nil, err
	}
	return client, nil
}
