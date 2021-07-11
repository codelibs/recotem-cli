package api

import (
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

func (c Client) GetToken(username string, password string) (string, error) {
	client, err := openapi.NewClientWithResponses(c.Url)
	if err != nil {
		return "", err
	}

	req := openapi.TokenCreateJSONRequestBody{}
	req.Username = username
	req.Password = password

	resp, err := client.TokenCreateWithResponse(c.Context, req)
	if err != nil {
		return "", err
	}

	return resp.JSON200.Token, nil
}
