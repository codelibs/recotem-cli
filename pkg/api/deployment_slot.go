package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetDeploymentSlots(page, pageSize, project *int) (*openapi.PaginatedDeploymentSlotList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.DeploymentSlotListParams
	if page != nil || pageSize != nil || project != nil {
		req = openapi.DeploymentSlotListParams{
			Page:     page,
			PageSize: pageSize,
			Project:  project,
		}
	}
	resp, err := client.DeploymentSlotListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) CreateDeploymentSlot(name string, project int, trainedModel *int, isActive bool) (*openapi.DeploymentSlot, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.DeploymentSlotCreateJSONRequestBody{
		Name:         name,
		Project:      project,
		TrainedModel: trainedModel,
		IsActive:     isActive,
	}
	resp, err := client.DeploymentSlotCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetDeploymentSlot(id int) (*openapi.DeploymentSlot, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.DeploymentSlotRetrieveWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) UpdateDeploymentSlot(id int, name *string, trainedModel *int) (*openapi.DeploymentSlot, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.DeploymentSlotUpdateJSONRequestBody{
		Name:         name,
		TrainedModel: trainedModel,
	}
	resp, err := client.DeploymentSlotUpdateWithResponse(c.Context, id, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteDeploymentSlot(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.DeploymentSlotDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
