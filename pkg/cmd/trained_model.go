package cmd

import (
	"fmt"
	"strconv"

	"github.com/urfave/cli/v2"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/utils"
)

func TrainedModelCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "trained-model",
		Aliases: []string{"tm"},
		Usage:   "options for trained model",
		Subcommands: []*cli.Command{
			trainedModelCreateCommand(),
			trainedModelDownloadCommand(),
			trainedModelListCommand(),
		},
	}
	return &cmd
}

func trainedModelCreateCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "create",
		Usage: "create a new trained model",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "configuration",
				Aliases:  []string{"c"},
				Usage:    "Configuration",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "data-loc",
				Aliases:  []string{"dl"},
				Usage:    "Data loc",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "file",
				Aliases: []string{"f"},
				Usage:   "File",
			},
			&cli.StringFlag{
				Name:    "irspack-version",
				Aliases: []string{"iv"},
				Usage:   "irspack version",
			},
		},
		Action: func(c *cli.Context) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(c.Context, config)
			configuration, err := strconv.Atoi(c.String("configuration"))
			if err != nil {
				return err
			}
			dataLoc, err := strconv.Atoi(c.String("data-loc"))
			if err != nil {
				return err
			}
			trainedModel, err := client.CreateTrainedModel(
				configuration,
				dataLoc,
				utils.NilOrString(c.String("file")),
				utils.NilOrString(c.String("irspack-version")))
			if err != nil {
				return err
			}
			fmt.Println("Created Trained Model ID: ", trainedModel.Id)
			return nil
		},
	}
	return &cmd
}

func trainedModelListCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: "get trained models",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "data-loc",
				Aliases: []string{"dl"},
				Usage:   "Data loc",
			},
			&cli.StringFlag{
				Name:    "data-loc-project",
				Aliases: []string{"dlp"},
				Usage:   "Data loc project",
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
			tmList, err := client.GetTrainedModels(
				utils.NilOrInt(c.String("data-loc")),
				utils.NilOrInt(c.String("data-loc-project")),
				utils.NilOrInt(c.String("id")),
				utils.NilOrInt(c.String("page")),
				utils.NilOrInt(c.String("page-size")),
			)
			if err != nil {
				return err
			}
			for _, x := range *tmList.Results {
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
	}
	return &cmd
}

func trainedModelDownloadCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "download",
		Usage: "download trained model",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"d"},
				Usage:    "Trained model ID",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "output",
				Aliases:  []string{"o"},
				Usage:    "Output filename",
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
			outputFilename := c.String("output")
			err = client.DownloadTrainedModel(id, outputFilename)
			if err != nil {
				return err
			}
			fmt.Println(outputFilename)
			return nil
		},
	}
	return &cmd
}
