package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

type TokenRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string
}

func NewTokenRequestBody(username string, password string) (body TokenRequestBody) {
	body = TokenRequestBody{}
	body.Username = username
	body.Password = password
	return
}

func (c Client) GetToken(username string, password string) (*openapi.AuthToken, error) {
	client, err := openapi.NewClientWithResponses(c.Config.Url)
	if err != nil {
		return nil, err
	}

	req := openapi.TokenCreateJSONRequestBody{}
	req.Username = username
	req.Password = password

	resp, err := client.TokenCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}
