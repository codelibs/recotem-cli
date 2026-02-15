package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newPingCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "ping",
		Short: "Check server health",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.Ping()
			if err != nil {
				return err
			}
			fmt.Println(string(result))
			return nil
		},
	}
}
