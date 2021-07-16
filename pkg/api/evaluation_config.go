package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateEvaluationConfig(name *string, cutoff *int, targetMetric *openapi.TargetMetricEnum) (*openapi.EvaluationConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.EvaluationConfigCreateJSONRequestBody{}
	req.Name = name
	req.Cutoff = cutoff
	req.TargetMetric = targetMetric
	resp, err := client.EvaluationConfigCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) GetEvaluationConfigs(id *int, name *string, unnamed *bool) (*[]openapi.EvaluationConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.EvaluationConfigListParams
	if id != nil || name != nil || unnamed != nil {
		req = openapi.EvaluationConfigListParams{}
		req.Id = id
		req.Name = name
		req.Unnamed = unnamed
	}
	resp, err := client.EvaluationConfigListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}
