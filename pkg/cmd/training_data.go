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

func TrainingDataCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "training-data",
		Aliases: []string{"td"},
		Usage:   "options for training data",
		Subcommands: []*cli.Command{
			trainingDataDeleteCommand(),
			trainingDataListCommand(),
			trainingDataUploadCommand(),
		},
	}
	return &cmd
}

func trainingDataUploadCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "upload",
		Usage: "upload a new training data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "project",
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
			id, err := strconv.Atoi(c.String("project"))
			if err != nil {
				return err
			}
			trainingData, err := client.UploadTrainingData(id, c.String("file"))
			if err != nil {
				return err
			}
			printTrainingData(*trainingData)
			return nil
		},
	}
	return &cmd
}

func trainingDataDeleteCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "delete",
		Usage: "delete the training data",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "id",
				Aliases:  []string{"i"},
				Usage:    "Training data ID",
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
			err = client.DeleteTrainingData(id)
			if err != nil {
				return err
			}
			fmt.Println(id)
			return nil
		},
	}
	return &cmd
}

func trainingDataListCommand() *cli.Command {
	cmd := cli.Command{
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
				printTrainingData(x)
			}
			return nil
		},
	}
	return &cmd
}

func printTrainingData(x openapi.TrainingData) {
	fmt.Println(x.Id,
		x.Project,
		utils.Atoa(x.Basename),
		x.Filesize,
		utils.FormatTime(x.InsDatetime))
}
