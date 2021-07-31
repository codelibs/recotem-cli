package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/openapi"
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
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"fmt"},
				Usage:       "Output format",
				DefaultText: "line",
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
			project, err := client.CreateProject(name, userColumn, itemColumn,
				utils.NilOrString(c.String("time-column")))
			if err != nil {
				return err
			}
			printProject(c.String("format"), *project)
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
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"fmt"},
				Usage:       "Output format",
				DefaultText: "line",
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
			utils.PrintId(c.String("format"), id)
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
			&cli.StringFlag{
				Name:        "format",
				Aliases:     []string{"fmt"},
				Usage:       "Output format",
				DefaultText: "line",
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
				printProject(c.String("format"), x)
			}
			return nil
		},
	}
	return &cmd
}

func printProject(format string, x openapi.Project) {
	if format == "json" {
		body := map[string]string{
			"id":          strconv.Itoa(x.Id),
			"name":        x.Name,
			"user_column": x.UserColumn,
			"item_column": x.ItemColumn,
			"time_column": utils.Atoa(x.TimeColumn)}
		bytes, err := json.Marshal(body)
		if err != nil {
			fmt.Println("JSON marshal error: ", err)
			return
		}
		fmt.Println(string(bytes))
	} else {
		fmt.Println(x.Id,
			x.Name,
			x.UserColumn,
			x.ItemColumn,
			utils.Atoa(x.TimeColumn))
	}
}
