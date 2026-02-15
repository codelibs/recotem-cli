package api

import (
	"fmt"

	openapi_types "github.com/oapi-codegen/runtime/types"
	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) GetUsers(page, pageSize *int) (*openapi.PaginatedUserList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.UserListParams
	if page != nil || pageSize != nil {
		req = openapi.UserListParams{
			Page:     page,
			PageSize: pageSize,
		}
	}
	resp, err := client.UserListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) CreateUser(username, password, email string) (*openapi.User, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	emailVal := openapi_types.Email(email)
	req := openapi.UserCreateJSONRequestBody{
		Username: username,
		Password: &password,
		Email:    emailVal,
	}
	resp, err := client.UserCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetUser(id int) (*openapi.User, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.UserRetrieveWithResponse(c.Context, id)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) UpdateUser(id int, email *string, isActive *bool) (*openapi.User, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.UserUpdateJSONRequestBody{
		IsActive: isActive,
	}
	if email != nil {
		emailVal := openapi_types.Email(*email)
		req.Email = &emailVal
	}
	resp, err := client.UserUpdateWithResponse(c.Context, id, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteUser(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.UserDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeactivateUser(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.UserDeactivateWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) ActivateUser(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.UserActivateWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) ResetUserPassword(id int, newPassword string) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	req := openapi.UserResetPasswordJSONRequestBody{
		NewPassword: newPassword,
	}
	resp, err := client.UserResetPasswordWithResponse(c.Context, id, req)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
