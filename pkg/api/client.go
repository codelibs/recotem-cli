package api

import (
	"context"

	"recotem.org/cli/recotem/pkg/cfg"
)

type Client struct {
	Context context.Context
	Url     string
}

func NewClient(context context.Context, config cfg.RecotemConfig) Client {
	client := Client{}
	client.Context = context
	client.Url = config.Url
	return client
}
