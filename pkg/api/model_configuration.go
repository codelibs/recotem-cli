package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateModelConfiguration(name *string, project int, recommenderClassName string,
	parametersJson string) (*openapi.ModelConfiguration, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.ModelConfigurationCreateJSONRequestBody{
		Name:                 name,
		Project:              project,
		RecommenderClassName: recommenderClassName,
		ParametersJson:       parametersJson,
	}
	resp, err := client.ModelConfigurationCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteModelConfiguration(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.ModelConfigurationDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetModelConfigurations(id *int, page *int, pageSize *int,
	project *int) (*openapi.PaginatedModelConfigurationList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.ModelConfigurationListParams
	if id != nil || page != nil || pageSize != nil || project != nil {
		req = openapi.ModelConfigurationListParams{
			Id:       id,
			Page:     page,
			PageSize: pageSize,
			Project:  project,
		}
	}
	resp, err := client.ModelConfigurationListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) UpdateModelConfiguration(id int, name *string,
	parametersJson *string) (*openapi.ModelConfiguration, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.ModelConfigurationUpdateJSONRequestBody{
		Name:           name,
		ParametersJson: parametersJson,
	}
	resp, err := client.ModelConfigurationUpdateWithResponse(c.Context, id, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
