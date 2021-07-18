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

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
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

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}
