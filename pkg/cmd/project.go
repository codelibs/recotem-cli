package cmd

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/utils"
)

func ProjectCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "project",
		Aliases: []string{"p"},
		Usage:   "tasks for a project",
		Subcommands: []*cli.Command{
			projectCreateCommand(),
			projectDeleteCommand(),
			projectListCommand(),
		},
	}
	return &cmd
}

func projectCreateCommand() *cli.Command {
	cmd := cli.Command{
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
	}
	return &cmd
}

func projectDeleteCommand() *cli.Command {
	cmd := cli.Command{
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
	}
	return &cmd
}

func projectListCommand() *cli.Command {
	cmd := cli.Command{
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
	}
	return &cmd
}
