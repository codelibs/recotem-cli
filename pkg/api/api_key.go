package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetApiKeys(page, pageSize *int) (*openapi.PaginatedApiKeyList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.ApiKeyListParams
	if page != nil || pageSize != nil {
		req = openapi.ApiKeyListParams{
			Page:     page,
			PageSize: pageSize,
		}
	}
	resp, err := client.ApiKeyListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) CreateApiKey(name string) (*openapi.ApiKey, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.ApiKeyCreateJSONRequestBody{
		Name: name,
	}
	resp, err := client.ApiKeyCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetApiKey(id int) (*openapi.ApiKey, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.ApiKeyRetrieveWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteApiKey(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.ApiKeyDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) RevokeApiKey(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.ApiKeyRevokeWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
