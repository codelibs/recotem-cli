package api

import "recotem.org/cli/recotem/pkg/cfg"

type Client struct {
	Url string
}

func NewClient(config cfg.RecotemConfig) Client {
	client := Client{}
	client.Url=config.Url
	return client
}
