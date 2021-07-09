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
				Action: func(c *cli.Context) error {
					username, password, err := utils.Credentials()
					if err != nil {
						return err
					}
					// username := "admin"
					// password := "very_bad_password"
					config, err := cfg.LoadRecotemConfig()
					if err != nil {
						return err
					}
					client:=api.NewClient(config)
					token, err := client.GetToken(username, password)
					if err != nil {
						return err
					}
					config.Token = token
					err = cfg.SaveRecotemConfig(config)
					if err != nil {
						return err
					}
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
