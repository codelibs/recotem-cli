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
					{
						Name:  "list",
						Usage: "get training data",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:    "id",
								Aliases: []string{"i"},
								Usage:   "Training data ID",
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
							tdList, err := client.GetTrainingData(
								utils.NilOrInt(c.String("id")),
								utils.NilOrInt(c.String("page")),
								utils.NilOrInt(c.String("page-size")),
								utils.NilOrInt(c.String("project")))
							if err != nil {
								return err
							}
							for _, x := range *tdList.Results {
								fmt.Println(x.Id, *x.Basename, x.Filesize, x.InsDatetime)
							}
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
			{
				Name:    "evaluation-config",
				Aliases: []string{"ec"},
				Usage:   "options for evaluation config",
				Subcommands: []*cli.Command{
					{
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
					},
					{
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
					},
				},
			},
			{
				Name:    "parameter-tuning-job",
				Aliases: []string{"ptj"},
				Usage:   "options for parameter tuning job",
				Subcommands: []*cli.Command{
					{
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
							fmt.Println("Created Parameter Tuning Job ID: ", parameterTuningJob.Id)
							return nil
						},
					},
					{
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
								Usage:   "Prameter tuning job ID",
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
								size := len(x.TaskLinks)
								if size > 0 {
									task := x.TaskLinks[size-1].Task
									fmt.Println(x.Id, x.InsDatetime, *task.Status)
								} else {
									fmt.Println(x.Id)
								}
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
