package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetRetrainingSchedules(page, pageSize, deploymentSlot *int) (*openapi.PaginatedRetrainingScheduleList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.RetrainingScheduleListParams
	if page != nil || pageSize != nil || deploymentSlot != nil {
		req = openapi.RetrainingScheduleListParams{
			Page:           page,
			PageSize:       pageSize,
			DeploymentSlot: deploymentSlot,
		}
	}
	resp, err := client.RetrainingScheduleListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) CreateRetrainingSchedule(deploymentSlot int, cronExpression string, isActive bool) (*openapi.RetrainingSchedule, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.RetrainingScheduleCreateJSONRequestBody{
		DeploymentSlot: deploymentSlot,
		CronExpression: cronExpression,
		IsActive:       isActive,
	}
	resp, err := client.RetrainingScheduleCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetRetrainingSchedule(id int) (*openapi.RetrainingSchedule, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.RetrainingScheduleRetrieveWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) UpdateRetrainingSchedule(id int, cronExpression *string, isActive *bool) (*openapi.RetrainingSchedule, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.RetrainingScheduleUpdateJSONRequestBody{
		CronExpression: cronExpression,
		IsActive:       isActive,
	}
	resp, err := client.RetrainingScheduleUpdateWithResponse(c.Context, id, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteRetrainingSchedule(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.RetrainingScheduleDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) TriggerRetrainingSchedule(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.RetrainingScheduleTriggerWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
