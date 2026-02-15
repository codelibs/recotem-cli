package api

import (
	"fmt"
	"time"

	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/openapi"
)

// LoginResult contains the tokens from a successful login
type LoginResult struct {
	AccessToken  string
	RefreshToken string
	ExpiresAt    time.Time
}

// Login authenticates with username and password, returns JWT tokens
func (c Client) Login(username, password string) (*LoginResult, error) {
	client, err := c.newUnauthenticatedClient()
	if err != nil {
		return nil, err
	}

	req := openapi.AuthLoginJSONRequestBody{
		Username: &username,
		Password: password,
	}

	resp, err := client.AuthLoginWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		// Default expiry: 5 minutes from now (typical JWT access token lifetime)
		expiresAt := time.Now().Add(5 * time.Minute)
		return &LoginResult{
			AccessToken:  resp.JSON200.AccessToken,
			RefreshToken: resp.JSON200.RefreshToken,
			ExpiresAt:    expiresAt,
		}, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

// Logout blacklists the refresh token
func (c Client) Logout() error {
	if c.Config.RefreshToken == "" {
		return nil
	}

	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	req := openapi.AuthTokenBlacklistJSONRequestBody{
		Refresh: c.Config.RefreshToken,
	}

	resp, err := client.AuthTokenBlacklistWithResponse(c.Context, req)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

// RefreshToken refreshes the access token using the refresh token
func (c *Client) RefreshToken() (*LoginResult, error) {
	if c.Config.RefreshToken == "" {
		return nil, fmt.Errorf("no refresh token available, please login again")
	}

	client, err := c.newUnauthenticatedClient()
	if err != nil {
		return nil, err
	}

	req := openapi.AuthTokenRefreshJSONRequestBody{
		Refresh: &c.Config.RefreshToken,
	}

	resp, err := client.AuthTokenRefreshWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		expiresAt := time.Now().Add(5 * time.Minute)
		result := &LoginResult{
			AccessToken:  *resp.JSON200.Access,
			RefreshToken: c.Config.RefreshToken,
			ExpiresAt:    expiresAt,
		}
		// Update client's config
		c.Config.AccessToken = result.AccessToken
		c.Config.ExpiresAt = &expiresAt
		return result, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

// EnsureValidToken checks if the token is expired and refreshes if needed.
// Saves updated config if token was refreshed.
func (c *Client) EnsureValidToken() error {
	// Skip if using API key or legacy token
	if c.Config.ApiKey != "" || c.Config.Token != "" {
		return nil
	}

	if c.Config.AccessToken == "" {
		return fmt.Errorf("not authenticated, please login first")
	}

	if !c.Config.IsTokenExpired() {
		return nil
	}

	result, err := c.RefreshToken()
	if err != nil {
		return fmt.Errorf("token refresh failed: %w (please login again)", err)
	}

	c.Config.AccessToken = result.AccessToken
	c.Config.ExpiresAt = &result.ExpiresAt
	return cfg.SaveRecotemConfig(c.Config)
}
