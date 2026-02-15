package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetTaskLogs(page, pageSize *int, task *int) ([]byte, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.TaskLogListParams
	if page != nil || pageSize != nil || task != nil {
		req = openapi.TaskLogListParams{
			Page:     page,
			PageSize: pageSize,
			Task:     task,
		}
	}
	resp, err := client.TaskLogListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return resp.Body, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
