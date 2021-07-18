package cmd

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func ModelConfigurationCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "model-configuration",
		Aliases: []string{"mc"},
		Usage:   "options for model configuration",
		Subcommands: []*cli.Command{
			modelConfigurationCreateCommand(),
			modelConfigurationListCommand(),
		},
	}
	return &cmd
}

func modelConfigurationCreateCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "create",
		Usage: "create a new model configuration",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Name",
			},
			&cli.StringFlag{
				Name:     "project",
				Aliases:  []string{"p"},
				Usage:    "Project",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "recommender-class-name",
				Aliases:  []string{"rcn"},
				Usage:    "Recommender class name",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "parameters-json",
				Aliases:  []string{"pj"},
				Usage:    "Parameters json",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(c.Context, config)
			project, err := strconv.Atoi(c.String("project"))
			if err != nil {
				return err
			}
			modelConfiguration, err := client.CreateModelConfiguration(
				utils.NilOrString(c.String("name")),
				project,
				c.String("recommender-class-name"),
				c.String("parameters-json"))
			if err != nil {
				return err
			}
			printModelConfiguration(*modelConfiguration)
			return nil
		},
	}
	return &cmd
}

func modelConfigurationListCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: "get model configurations",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "id",
				Aliases: []string{"i"},
				Usage:   "Model configuration ID",
			},
			&cli.StringFlag{
				Name:    "page",
				Aliases: []string{"p"},
				Usage:   "Page",
			},
			&cli.StringFlag{
				Name:    "page-size",
				Aliases: []string{"ps"},
				Usage:   "Page size",
			},
			&cli.StringFlag{
				Name:    "project",
				Aliases: []string{"pj"},
				Usage:   "Project",
			},
		},
		Action: func(c *cli.Context) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(c.Context, config)
			modelConfigurations, err := client.GetModelConfigurations(
				utils.NilOrInt(c.String("id")),
				utils.NilOrInt(c.String("page")),
				utils.NilOrInt(c.String("page-size")),
				utils.NilOrInt(c.String("project")))
			if err != nil {
				return err
			}
			for _, x := range *modelConfigurations.Results {
				printModelConfiguration(x)
			}
			return nil
		},
	}
	return &cmd
}

func printModelConfiguration(x openapi.ModelConfiguration) {
	fmt.Println(x.Id, x.Project, x.RecommenderClassName, x.TuningJob, *x.Name)
}
