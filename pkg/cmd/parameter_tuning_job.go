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

func ParameterTuningJobCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "parameter-tuning-job",
		Aliases: []string{"ptj"},
		Usage:   "options for parameter tuning job",
		Subcommands: []*cli.Command{
			parameterTuningJobCreateCommand(),
			parameterTuningJobDeleteCommand(),
			parameterTuningJobListCommand(),
		},
	}
	return &cmd
}

func parameterTuningJobCreateCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "create",
		Usage: "create a new parameter tuning job",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "data",
				Aliases:  []string{"d"},
				Usage:    "Data",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "split",
				Aliases:  []string{"s"},
				Usage:    "Split",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "evaluation",
				Aliases:  []string{"e"},
				Usage:    "Evaluation",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "n-tasks-parallel",
				Aliases: []string{"ntp"},
				Usage:   "N tasks parallel",
			},
			&cli.StringFlag{
				Name:    "n-trials",
				Aliases: []string{"nt"},
				Usage:   "N trials",
			},
			&cli.StringFlag{
				Name:    "memory-budget",
				Aliases: []string{"mb"},
				Usage:   "Memory budget",
			},
			&cli.StringFlag{
				Name:    "timeout-overall",
				Aliases: []string{"to"},
				Usage:   "Timeout overall",
			},
			&cli.StringFlag{
				Name:    "timeout-singlestep",
				Aliases: []string{"ts"},
				Usage:   "Timeout singlestep",
			},
			&cli.StringFlag{
				Name:    "random-seed",
				Aliases: []string{"rs"},
				Usage:   "Random seed",
			},
			&cli.StringFlag{
				Name:    "tried-algorithm-json",
				Aliases: []string{"taj"},
				Usage:   "Tried algorithm json",
			},
			&cli.StringFlag{
				Name:    "irspack-version",
				Aliases: []string{"iv"},
				Usage:   "irspack version",
			},
			&cli.StringFlag{
				Name:    "train-after-tuning",
				Aliases: []string{"tat"},
				Usage:   "Train after tuning",
			},
			&cli.StringFlag{
				Name:    "best-score",
				Aliases: []string{"bs"},
				Usage:   "Best score",
			},
			&cli.StringFlag{
				Name:    "tuned-model",
				Aliases: []string{"tm"},
				Usage:   "Tuned model",
			},
			&cli.StringFlag{
				Name:    "best-config",
				Aliases: []string{"bc"},
				Usage:   "Best config",
			},
		},
		Action: func(c *cli.Context) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(c.Context, config)
			data, err := strconv.Atoi(c.String("data"))
			if err != nil {
				return err
			}
			split, err := strconv.Atoi(c.String("split"))
			if err != nil {
				return err
			}
			evaluation, err := strconv.Atoi(c.String("evaluation"))
			if err != nil {
				return err
			}
			parameterTuningJob, err := client.CreateParameterTuningJob(
				data,
				split,
				evaluation,
				utils.NilOrInt(c.String("n-tasks-parallel")),
				utils.NilOrInt(c.String("n-trials")),
				utils.NilOrInt(c.String("memory-budget")),
				utils.NilOrInt(c.String("timeout-overall")),
				utils.NilOrInt(c.String("timeout-singlestep")),
				utils.NilOrInt(c.String("random-seed")),
				utils.NilOrString(c.String("tried-algorithm-json")),
				utils.NilOrString(c.String("irspack-version")),
				utils.NilOrBool(c.String("train-after-tuning")),
				utils.NilOrFloat32(c.String("best-score")),
				utils.NilOrInt(c.String("tuned-model")),
				utils.NilOrInt(c.String("best-config")))
			if err != nil {
				return err
			}
			printParameterTuningJob(*parameterTuningJob)
			return nil
		},
	}
	return &cmd
}

func parameterTuningJobDeleteCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "delete",
		Usage: "delete the parameter tuning job",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Parameter tuning job ID",
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
			err = client.DeleteParameterTuningJob(id)
			if err != nil {
				return err
			}
			fmt.Println(id)
			return nil
		},
	}
	return &cmd
}

func parameterTuningJobListCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: "get parameter tuning jobs",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "data",
				Aliases: []string{"d"},
				Usage:   "Data",
			},
			&cli.StringFlag{
				Name:    "data-project",
				Aliases: []string{"dp"},
				Usage:   "Data project",
			},
			&cli.StringFlag{
				Name:    "id",
				Aliases: []string{"i"},
				Usage:   "Parameter tuning job ID",
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
		},
		Action: func(c *cli.Context) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(c.Context, config)
			ptjList, err := client.GetParameterTuningJobs(
				utils.NilOrInt(c.String("data")),
				utils.NilOrInt(c.String("data-project")),
				utils.NilOrInt(c.String("id")),
				utils.NilOrInt(c.String("page")),
				utils.NilOrInt(c.String("page-size")),
			)
			if err != nil {
				return err
			}
			for _, x := range *ptjList.Results {
				printParameterTuningJob(x)
			}
			return nil
		},
	}
	return &cmd
}

func printParameterTuningJob(x openapi.ParameterTuningJob) {
	size := len(x.TaskLinks)
	if size > 0 {
		task := x.TaskLinks[size-1].Task
		var status string
		if task.Status != nil {
			status = string(*task.Status)
		} else {
			status = utils.NoValue
		}
		fmt.Println(x.Id,
			utils.FormatTime(x.InsDatetime),
			status,
			utils.Itoa(x.TunedModel))
	} else {
		fmt.Println(x.Id)
	}
}
