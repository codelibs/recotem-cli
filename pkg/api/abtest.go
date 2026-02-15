package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetAbTests(page, pageSize, project *int) (*openapi.PaginatedAbTestList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.AbTestListParams
	if page != nil || pageSize != nil || project != nil {
		req = openapi.AbTestListParams{
			Page:     page,
			PageSize: pageSize,
			Project:  project,
		}
	}
	resp, err := client.AbTestListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) CreateAbTest(name string, project int, slots []int) (*openapi.AbTest, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.AbTestCreateJSONRequestBody{
		Name:    name,
		Project: project,
		Slots:   slots,
	}
	resp, err := client.AbTestCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetAbTest(id int) (*openapi.AbTest, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.AbTestRetrieveWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) UpdateAbTest(id int, name *string, slots *[]int) (*openapi.AbTest, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.AbTestUpdateJSONRequestBody{
		Name:  name,
		Slots: slots,
	}
	resp, err := client.AbTestUpdateWithResponse(c.Context, id, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteAbTest(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.AbTestDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) StartAbTest(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.AbTestStartWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) StopAbTest(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.AbTestStopWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetAbTestResults(id int) (*[]openapi.AbTestResult, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.AbTestResultsWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) PromoteAbTestWinner(id int, slotId int) (*openapi.AbTest, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.AbTestPromoteWinnerJSONRequestBody{
		SlotId: slotId,
	}
	resp, err := client.AbTestPromoteWinnerWithResponse(c.Context, id, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
