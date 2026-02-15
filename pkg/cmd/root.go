package cmd

import (
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"

	"github.com/spf13/cobra"
)

var (
	outputFormat string
	apiKeyFlag   string
)

func NewRootCmd(version, commit, buildTime string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "recotem",
		Short:        "CLI for recotem recommendation system",
		Long:         "Command line interface for managing recotem recommendation system resources.",
		SilenceUsage: true,
	}

	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "text", "Output format (text, json, yaml)")
	rootCmd.PersistentFlags().StringVar(&apiKeyFlag, "api-key", "", "API key for authentication")

	rootCmd.AddCommand(
		newLoginCmd(),
		newLogoutCmd(),
		newVersionCmd(version, commit, buildTime),
		newCompletionCmd(),
		newPingCmd(),
		newProjectCmd(),
		newTrainingDataCmd(),
		newItemMetaDataCmd(),
		newTrainedModelCmd(),
		newModelConfigurationCmd(),
		newEvaluationConfigCmd(),
		newSplitConfigCmd(),
		newParameterTuningJobCmd(),
		newApiKeyCmd(),
		newDeploymentSlotCmd(),
		newAbTestCmd(),
		newConversionEventCmd(),
		newRetrainingScheduleCmd(),
		newRetrainingRunCmd(),
		newTaskLogCmd(),
		newUserCmd(),
	)

	return rootCmd
}

func newClientFromCmd(cmd *cobra.Command) (api.Client, error) {
	config, err := cfg.LoadRecotemConfig()
	if err != nil {
		return api.Client{}, err
	}

	// Override API key from flag if provided
	if apiKeyFlag != "" {
		config.ApiKey = apiKeyFlag
	}

	client := api.NewClient(cmd.Context(), config)
	if err := client.EnsureValidToken(); err != nil {
		return api.Client{}, err
	}
	return client, nil
}

func getOutputFormat() string {
	if outputFormat == "" {
		return "text"
	}
	return outputFormat
}
