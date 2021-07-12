package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/utils"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "login",
				Aliases: []string{},
				Usage:   "get an gaccess token",
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
					config.Token = token
					err = cfg.SaveRecotemConfig(config)
					if err != nil {
						return err
					}
					fmt.Println("Updated your token.")
					return nil
				},
			},
			{
				Name:    "complete",
				Aliases: []string{"c"},
				Usage:   "complete a task on the list",
				Action: func(c *cli.Context) error {
					fmt.Println("completed task: ", c.Args().First())
					return nil
				},
			},
			{
				Name:    "template",
				Aliases: []string{"t"},
				Usage:   "options for task templates",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add a new template",
						Action: func(c *cli.Context) error {
							fmt.Println("new task template: ", c.Args().First())
							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(c *cli.Context) error {
							fmt.Println("removed task template: ", c.Args().First())
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
