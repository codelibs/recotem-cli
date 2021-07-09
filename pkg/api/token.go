package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	requestBody := NewTokenRequestBody(username, password)
	jsonRequestBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	url := c.Url + "/api/token/"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonRequestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Could not access Recotem server. HTTP status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	responseBody := TokenResponse{}
	err = json.Unmarshal([]byte(body), &responseBody)
	if err != nil {
		return "", err
	}

	return responseBody.Token, nil
}
