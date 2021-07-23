package api

import (
	"fmt"
	"io/ioutil"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateTrainedModel(configuration int, dataLoc int, file *string,
	irspackVersion *string) (*openapi.TrainedModel, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.TrainedModelCreateJSONRequestBody{
		Configuration:  configuration,
		DataLoc:        dataLoc,
		File:           file,
		IrspackVersion: irspackVersion,
	}
	resp, err := client.TrainedModelCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) DeleteTrainedModel(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.TrainedModelDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) GetTrainedModels(dataLoc *int, dataLocProject *int, id *int, page *int,
	pageSize *int) (*openapi.PaginatedTrainedModelList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.TrainedModelListParams
	if dataLoc != nil || dataLocProject != nil || id != nil || page != nil || pageSize != nil {
		req = openapi.TrainedModelListParams{
			DataLoc:        dataLoc,
			DataLocProject: dataLocProject,
			Id:             id,
			Page:           page,
			PageSize:       pageSize,
		}
	}
	resp, err := client.TrainedModelListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) DownloadTrainedModel(id int, output string) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.TrainedModelDownloadFileRetrieveWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
	}

	err = ioutil.WriteFile(output, resp.Body, 0644)
	if err != nil {
		return err
	}

	return nil
}
