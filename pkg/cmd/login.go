package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/utils"
)

func LoginCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "login",
		Aliases: []string{},
		Usage:   "get an access token",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "username",
				Aliases: []string{"u"},
				Usage:   "Username",
			},
			&cli.StringFlag{
				Name:    "password",
				Aliases: []string{"p"},
				Usage:   "Password",
			},
		},
		Action: func(c *cli.Context) error {
			username, password, err := utils.Credentials(c.String("username"), c.String("password"))
			if err != nil {
				return err
			}
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(c.Context, config)
			token, err := client.GetToken(username, password)
			if err != nil {
				return err
			}
			config.Token = token.Token
			err = cfg.SaveRecotemConfig(config)
			if err != nil {
				return err
			}
			fmt.Println("Updated your token.")
			return nil
		},
	}
	return &cmd
}
