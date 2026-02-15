package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetConversionEvents(page, pageSize *int, abTest *int) (*openapi.PaginatedConversionEventList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.ConversionEventListParams
	if page != nil || pageSize != nil || abTest != nil {
		req = openapi.ConversionEventListParams{
			Page:     page,
			PageSize: pageSize,
			AbTest:   abTest,
		}
	}
	resp, err := client.ConversionEventListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) CreateConversionEvent(abTest int, userId string, itemId *string, eventType string, slot int) (*openapi.ConversionEvent, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.ConversionEventCreateJSONRequestBody{
		AbTest:    abTest,
		UserId:    userId,
		ItemId:    itemId,
		EventType: eventType,
		Slot:      slot,
	}
	resp, err := client.ConversionEventCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetConversionEvent(id int) (*openapi.ConversionEvent, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.ConversionEventRetrieveWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) BatchCreateConversionEvents(events []openapi.ConversionEventCreate) ([]byte, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.ConversionEventBatchCreateJSONRequestBody{
		Events: events,
	}
	resp, err := client.ConversionEventBatchCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return resp.Body, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
