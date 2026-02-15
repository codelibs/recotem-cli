package api

import "fmt"

func (c Client) Ping() ([]byte, error) {
	client, err := c.newUnauthenticatedClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.PingWithResponse(c.Context)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return resp.Body, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
