package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateSplitConfig(name *string, scheme *openapi.SchemeEnum, heldoutRatio *float32,
	nHeldout *int, testUserRatio *float32, nTestUsers *int, randomSeed *int) (*openapi.SplitConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.SplitConfigCreateJSONRequestBody{
		Name:          name,
		Scheme:        scheme,
		HeldoutRatio:  heldoutRatio,
		NHeldout:      nHeldout,
		TestUserRatio: testUserRatio,
		NTestUsers:    nTestUsers,
		RandomSeed:    randomSeed,
	}
	resp, err := client.SplitConfigCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteSplitConfig(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.SplitConfigDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetSplitConfigs(id *int, name *string, unnamed *bool) (*[]openapi.SplitConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.SplitConfigListParams
	if id != nil || name != nil || unnamed != nil {
		req = openapi.SplitConfigListParams{
			Id:      id,
			Name:    name,
			Unnamed: unnamed,
		}
	}
	resp, err := client.SplitConfigListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
