package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/utils"
)

func EvaluationConfigCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "evaluation-config",
		Aliases: []string{"ec"},
		Usage:   "options for evaluation config",
		Subcommands: []*cli.Command{
			evaluationConfigCreateCommand(),
			evaluationConfigListCommand(),
		},
	}
	return &cmd
}

func evaluationConfigCreateCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "create",
		Usage: "create a new evaluation config",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "Name",
			},
			&cli.StringFlag{
				Name:    "cutoff",
				Aliases: []string{"c"},
				Usage:   "Cutoff",
			},
			&cli.StringFlag{
				Name:    "target-metric",
				Aliases: []string{"tm"},
				Usage:   "Target metric (ndcg|map|recall|hit)",
			},
		},
		Action: func(c *cli.Context) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(c.Context, config)
			evaluationConfig, err := client.CreateEvaluationConfig(
				utils.NilOrString(c.String("name")),
				utils.NilOrInt(c.String("cutoff")),
				utils.NilOrTargetMetric(c.String("target-metric")))
			if err != nil {
				return err
			}
			fmt.Println("Created Evaluation Config ID: ", evaluationConfig.Id)
			return nil
		},
	}
	return &cmd
}

func evaluationConfigListCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: "get evaluation configs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "id",
				Aliases: []string{"i"},
				Usage:   "Evaluation config ID",
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
			splitConfigs, err := client.GetEvaluationConfigs(
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
	}
	return &cmd
}
