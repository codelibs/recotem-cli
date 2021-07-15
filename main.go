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
						Name:  "list",
						Usage: "get projects",
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
							projects, err := client.GetProjects(
								utils.NilOrInt(c.String("id")),
								utils.NilOrString(c.String("name")))
							if err != nil {
								return err
							}
							for _, x := range *projects {
								fmt.Println(x.Id, x.Name)
							}
							return nil
						},
					},
				},
			},
			{
				Name:    "training-data",
				Aliases: []string{"td"},
				Usage:   "options for training data",
				Subcommands: []*cli.Command{
					{
						Name:  "upload",
						Usage: "upload a new training data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "project-id",
								Aliases:  []string{"p"},
								Usage:    "Project ID",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "file",
								Aliases:  []string{"f"},
								Usage:    "File for a training data",
								Required: true,
							},
						},
						Action: func(c *cli.Context) error {
							config, err := cfg.LoadRecotemConfig()
							if err != nil {
								return err
							}
							client := api.NewClient(c.Context, config)
							id, err := strconv.Atoi(c.String("project-id"))
							if err != nil {
								return err
							}
							data, err := client.UploadTrainingData(id, c.String("file"))
							if err != nil {
								return err
							}
							fmt.Println("Created Data ID: ", data.Id)
							return nil
						},
					},
				},
			},
			{
				Name:    "split-config",
				Aliases: []string{"sc"},
				Usage:   "options for split config",
				Subcommands: []*cli.Command{
					{
						Name:  "create",
						Usage: "create a new split config",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Usage:   "Name",
							},
							&cli.StringFlag{
								Name:    "scheme",
								Aliases: []string{"s"},
								Usage:   "Scheme (RG|TG|TU)",
							},
							&cli.StringFlag{
								Name:    "heldout-ratio",
								Aliases: []string{"hr"},
								Usage:   "Heldout ratio",
							},
							&cli.StringFlag{
								Name:    "n-heldout",
								Aliases: []string{"nh"},
								Usage:   "The number of heldout",
							},
							&cli.StringFlag{
								Name:    "test-user-ratio",
								Aliases: []string{"tur"},
								Usage:   "Test user ratio",
							},
							&cli.StringFlag{
								Name:    "n-test-users",
								Aliases: []string{"ntu"},
								Usage:   "The number of test users",
							},
							&cli.StringFlag{
								Name:    "random-seed",
								Aliases: []string{"rs"},
								Usage:   "Random seed",
							},
						},
						Action: func(c *cli.Context) error {
							config, err := cfg.LoadRecotemConfig()
							if err != nil {
								return err
							}
							client := api.NewClient(c.Context, config)
							splitConfig, err := client.CreateSplitConfig(
								utils.NilOrString(c.String("name")),
								utils.NilOrScheme(c.String("scheme")),
								utils.NilOrFloat32(c.String("heldout-ratio")),
								utils.NilOrInt(c.String("n-heldout")),
								utils.NilOrFloat32(c.String("test-user-ratio")),
								utils.NilOrInt(c.String("n-test-users")),
								utils.NilOrInt(c.String("random-seed")))
							if err != nil {
								return err
							}
							fmt.Println("Created Split Config ID: ", splitConfig.Id)
							return nil
						},
					},
					{
						Name:  "list",
						Usage: "get split configs",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "id",
								Aliases: []string{"i"},
								Usage:   "Split config ID",
							},
							&cli.StringFlag{
								Name:    "name",
								Aliases: []string{"n"},
								Usage:   "Name",
							},
							&cli.StringFlag{
								Name:    "unnamed",
								Aliases: []string{"u"},
								Usage:   "Unnamed",
							},
						},
						Action: func(c *cli.Context) error {
							config, err := cfg.LoadRecotemConfig()
							if err != nil {
								return err
							}
							client := api.NewClient(c.Context, config)
							splitConfigs, err := client.GetSplitConfigs(
								utils.NilOrInt(c.String("id")),
								utils.NilOrString(c.String("name")),
								utils.NilOrBool(c.String("unnamed")))
							if err != nil {
								return err
							}
							for _, x := range *splitConfigs {
								fmt.Println(x.Id, x.Name)
							}
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
