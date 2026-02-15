package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/api"
	"recotem.org/cli/recotem/pkg/cfg"
	"recotem.org/cli/recotem/pkg/utils"
)

func newLoginCmd() *cobra.Command {
	var username, password string

	cmd := &cobra.Command{
		Use:   "login",
		Short: "Authenticate with the recotem server",
		Long:  "Login to the recotem server using username and password to obtain JWT tokens.",
		RunE: func(cmd *cobra.Command, args []string) error {
			u, p, err := utils.Credentials(username, password)
			if err != nil {
				return err
			}
			config, err := cfg.LoadRecotemConfig()
			if err != nil {
				return err
			}
			client := api.NewClient(cmd.Context(), config)
			result, err := client.Login(u, p)
			if err != nil {
				return err
			}
			config.AccessToken = result.AccessToken
			config.RefreshToken = result.RefreshToken
			config.ExpiresAt = &result.ExpiresAt
			config.Token = "" // Clear legacy token
			err = cfg.SaveRecotemConfig(config)
			if err != nil {
				return err
			}
			fmt.Println("Login successful.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&username, "username", "u", "", "Username")
	cmd.Flags().StringVarP(&password, "password", "p", "", "Password")

	return cmd
}
