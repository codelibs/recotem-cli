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

func ItemMetaDataCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "item-meta-data",
		Aliases: []string{"imd"},
		Usage:   "options for item meta data",
		Subcommands: []*cli.Command{
			itemMetaDataListCommand(),
			itemMetaDataUploadCommand(),
		},
	}
	return &cmd
}

func itemMetaDataUploadCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "upload",
		Usage: "upload a new item meta data",
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
				Usage:    "File for a item-meta data",
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
			itemMetaData, err := client.UploadItemMetaData(id, c.String("file"))
			if err != nil {
				return err
			}
			printItemMetaData(*itemMetaData)
			return nil
		},
	}
	return &cmd
}

func itemMetaDataListCommand() *cli.Command {
	cmd := cli.Command{
		Name:  "list",
		Usage: "get item-meta data",
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
			tdList, err := client.GetItemMetaData(
				utils.NilOrInt(c.String("id")),
				utils.NilOrInt(c.String("page")),
				utils.NilOrInt(c.String("page-size")),
				utils.NilOrInt(c.String("project")))
			if err != nil {
				return err
			}
			for _, x := range *tdList.Results {
				printItemMetaData(x)
			}
			return nil
		},
	}
	return &cmd
}

func printItemMetaData(x openapi.ItemMetaData) {
	fmt.Println(x.Id, *x.Basename, x.Filesize, x.InsDatetime)
}
