package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func SplitConfigCommand() *cli.Command {
	cmd := cli.Command{
		Name:    "split-config",
		Aliases: []string{"sc"},
		Usage:   "options for split config",
		Subcommands: []*cli.Command{
			splitConfigCreateCommand(),
			splitConfigListCommand(),
		},
	}
	return &cmd
}

func splitConfigCreateCommand() *cli.Command {
	cmd := cli.Command{
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
			printSplitConfig(*splitConfig)
			return nil
		},
	}
	return &cmd
}

func splitConfigListCommand() *cli.Command {
	cmd := cli.Command{
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
				printSplitConfig(x)
			}
			return nil
		},
	}
	return &cmd
}

func printSplitConfig(x openapi.SplitConfig) {
	fmt.Println(x.Id,
		utils.Ftoa(x.HeldoutRatio),
		utils.Ftoa(x.TestUserRatio),
		utils.Itoa(x.RandomSeed))
}
