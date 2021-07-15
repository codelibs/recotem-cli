package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateSplitConfig(name *string, scheme *openapi.SchemeEnum, heldoutRatio *float32, nHeldout *int,
	testUserRatio *float32, nTestUsers *int, randomSeed *int) (*openapi.SplitConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.SplitConfigCreateJSONRequestBody{}
	req.Name = name
	req.Scheme = scheme
	req.HeldoutRatio = heldoutRatio
	req.NHeldout = nHeldout
	req.TestUserRatio = testUserRatio
	req.NTestUsers = nTestUsers
	req.RandomSeed = randomSeed
	resp, err := client.SplitConfigCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) GetSplitConfigs(id *int, name *string, unnamed *bool) (*[]openapi.SplitConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.SplitConfigListParams
	if id != nil || name != nil || unnamed != nil {
		req = openapi.SplitConfigListParams{}
		req.Id = id
		req.Name = name
		req.Unnamed = unnamed
	}
	resp, err := client.SplitConfigListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}
