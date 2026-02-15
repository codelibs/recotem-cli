package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func newVersionCmd(version, commit, buildTime string) *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("recotem version %s\n", version)
			fmt.Printf("  commit: %s\n", commit)
			fmt.Printf("  built:  %s\n", buildTime)
		},
	}
}
