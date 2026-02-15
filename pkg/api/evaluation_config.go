package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateEvaluationConfig(name *string, cutoff *int,
	targetMetric *openapi.TargetMetricEnum) (*openapi.EvaluationConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.EvaluationConfigCreateJSONRequestBody{
		Name:         name,
		Cutoff:       cutoff,
		TargetMetric: targetMetric,
	}
	resp, err := client.EvaluationConfigCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteEvaluationConfig(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.EvaluationConfigDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetEvaluationConfigs(id *int, name *string,
	unnamed *bool) (*[]openapi.EvaluationConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.EvaluationConfigListParams
	if id != nil || name != nil || unnamed != nil {
		req = openapi.EvaluationConfigListParams{
			Id:      id,
			Name:    name,
			Unnamed: unnamed,
		}
	}
	resp, err := client.EvaluationConfigListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) UpdateEvaluationConfig(id int, name *string, cutoff *int,
	targetMetric *openapi.TargetMetricEnum) (*openapi.EvaluationConfig, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.EvaluationConfigUpdateJSONRequestBody{
		Name:         name,
		Cutoff:       cutoff,
		TargetMetric: targetMetric,
	}
	resp, err := client.EvaluationConfigUpdateWithResponse(c.Context, id, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
