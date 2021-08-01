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

func EvaluationConfigCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "evaluation-config",
		Aliases: []string{"ec"},
		Usage:   "options for evaluation config",
		Subcommands: []*cli.Command{
			evaluationConfigCreateCommand(),
			evaluationConfigDeleteCommand(),
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
			evaluationConfig, err := client.CreateEvaluationConfig(
				utils.NilOrString(c.String("name")),
				utils.NilOrInt(c.String("cutoff")),
				utils.NilOrTargetMetric(c.String("target-metric")))
			if err != nil {
				return err
			}
			printEvaluationConfig(c.String("format"), *evaluationConfig)
			return nil
		},
	}
	return &cmd
}

func evaluationConfigDeleteCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "delete",
		Usage: "delete the evaluation config",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Evaluation config ID",
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
			err = client.DeleteEvaluationConfig(id)
			if err != nil {
				return err
			}
			utils.PrintId(c.String("format"), id)
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
			evaluationConfigs, err := client.GetEvaluationConfigs(
				utils.NilOrInt(c.String("id")),
				utils.NilOrString(c.String("name")),
				utils.NilOrBool(c.String("unnamed")))
			if err != nil {
				return err
			}
			for _, x := range *evaluationConfigs {
				printEvaluationConfig(c.String("format"), x)
			}
			return nil
		},
	}
	return &cmd
}

func printEvaluationConfig(format string, x openapi.EvaluationConfig) {
	var targetMetric string
	if x.TargetMetric != nil {
		targetMetric = string(*x.TargetMetric)
	} else {
		targetMetric = utils.NoValue
	}
	if format == "json" {
		body := map[string]string{
			"id":            strconv.Itoa(x.Id),
			"cutoff":        utils.Itoa(x.Cutoff),
			"target_metric": targetMetric,
			"name":          utils.Atoa(x.Name)}
		bytes, err := json.Marshal(body)
		if err != nil {
			fmt.Println("JSON marshal error: ", err)
			return
		}
		fmt.Println(string(bytes))
	} else {
		fmt.Println(
			x.Id,
			utils.Itoa(x.Cutoff),
			targetMetric,
			utils.FormatName(utils.Atoa(x.Name)))
	}
}
