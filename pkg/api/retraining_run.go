package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetRetrainingRuns(page, pageSize *int, schedule *int, status *openapi.RetrainingRunListParamsStatus) (*openapi.PaginatedRetrainingRunList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.RetrainingRunListParams
	if page != nil || pageSize != nil || schedule != nil || status != nil {
		req = openapi.RetrainingRunListParams{
			Page:     page,
			PageSize: pageSize,
			Schedule: schedule,
			Status:   status,
		}
	}
	resp, err := client.RetrainingRunListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetRetrainingRun(id int) (*openapi.RetrainingRun, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.RetrainingRunRetrieveWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
