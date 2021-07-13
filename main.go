package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
					config.Token = token.Token
					err = cfg.SaveRecotemConfig(config)
					if err != nil {
						return err
					}
					fmt.Println("Updated your token.")
					return nil
				},
			},
			{
				Name:    "project",
				Aliases: []string{"p"},
				Usage:   "tasks for a project",
				Subcommands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create a new project",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "name",
								Aliases:  []string{"n"},
								Usage:    "Project name",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "user-column",
								Aliases:  []string{"u"},
								Usage:    "User column",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "item-column",
								Aliases:  []string{"i"},
								Usage:    "Item column",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "time-column",
								Aliases: []string{"t"},
								Usage:   "Time column",
							},
						},
						Action: func(c *cli.Context) error {
							config, err := cfg.LoadRecotemConfig()
							if err != nil {
								return err
							}
							client := api.NewClient(c.Context, config)
							name := c.String("name")
							userColumn := c.String("user-column")
							itemColumn := c.String("item-column")
							timeColumn := c.String("time-column")
							project, err := client.CreateProject(name, userColumn, itemColumn, &timeColumn)
							if err != nil {
								return err
							}
							fmt.Println("Created Project ID: ", project.Id)
							return nil
						},
					},
					{
						Name:  "delete",
						Usage: "delete the project",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "id",
								Aliases:  []string{"i"},
								Usage:    "Project ID",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							config, err := cfg.LoadRecotemConfig()
							if err != nil {
								return err
							}
							client := api.NewClient(c.Context, config)
							id, err := strconv.Atoi(c.String("id"))
							if err != nil {
								return err
							}
							err = client.DeleteProject(id)
							if err != nil {
								return err
							}
							fmt.Println("Deleted Project ID: ", id)
							return nil
						},
					},
					{
						Name:  "get",
						Usage: "Get projects",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "id",
								Aliases: []string{"i"},
								Usage:   "Project ID",
							},
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Usage:   "Project name",
							},
						},
						Action: func(c *cli.Context) error {
							config, err := cfg.LoadRecotemConfig()
							if err != nil {
								return err
							}
							client := api.NewClient(c.Context, config)
							var id int
							idStr := c.String("id")
							if len(idStr) > 0 {
								id, err = strconv.Atoi(idStr)
								if err != nil {
									return err
								}
							}
							name := c.String("name")
							projects, err := client.GetProject(&id, &name)
							if err != nil {
								return err
							}
							for _, p := range *projects {
								fmt.Println(p.Id, p.Name)
							}
							return nil
						},
					},
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
