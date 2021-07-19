package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"recotem.org/cli/recotem/pkg/cmd"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			cmd.EvaluationConfigCommand(),
			cmd.ItemMetaDataCommand(),
			cmd.LoginCommand(),
			cmd.ModelConfigurationCommand(),
			cmd.ParameterTuningJobCommand(),
			cmd.ProjectCommand(),
			cmd.SplitConfigCommand(),
			cmd.TrainedModelCommand(),
			cmd.TrainingDataCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
