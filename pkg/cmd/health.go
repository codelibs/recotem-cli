package cmd

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"

	"github.com/spf13/cobra"
)

func newPingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Check server health",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(cmd.Context(), config)
			result, err := client.Ping()
			if err != nil {
				return err
			}
			fmt.Println(string(result))
			return nil
		},
	}
}
