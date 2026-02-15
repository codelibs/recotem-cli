package main

import (
	"os"

	"recotem.org/cli/recotem/pkg/cmd"
)

var (
	Version   = "dev"
	Commit    = "unknown"
	BuildTime = "unknown"
)

func main() {
	rootCmd := cmd.NewRootCmd(Version, Commit, BuildTime)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
