package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
)

func newLogoutCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "logout",
		Short: "Clear authentication tokens",
		Long:  "Logout from the recotem server by blacklisting the refresh token and clearing stored credentials.",
		RunE: func(cmd *cobra.Command, args []string) error {
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(cmd.Context(), config)
			// Try to blacklist the refresh token (best effort)
			if config.RefreshToken != "" {
				_ = client.Logout()
			}
			config.ClearTokens()
			err = cfg.SaveRecotemConfig(config)
			if err != nil {
				return err
			}
			fmt.Println("Logged out successfully.")
			return nil
		},
	}
}
