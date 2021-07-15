package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateProject(name string, userColumn string, itemColumn string, timeColumn *string) (*openapi.Project, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.ProjectCreateJSONRequestBody{}
	req.Name = name
	req.UserColumn = userColumn
	req.ItemColumn = itemColumn
	if len(*timeColumn) > 0 {
		req.TimeColumn = timeColumn
	}

	resp, err := client.ProjectCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) DeleteProject(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.ProjectDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) GetProjects(id *int, name *string) (*[]openapi.Project, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.ProjectListParams
	if id != nil || name != nil {
		req = openapi.ProjectListParams{}
		req.Id = id
		req.Name = name
	}
	resp, err := client.ProjectListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}
