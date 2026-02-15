package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/utils"
)

func newTaskLogCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "task-log",
		Aliases: []string{"tl"},
		Short:   "Manage task logs",
	}

	cmd.AddCommand(
		newTaskLogListCmd(),
	)

	return cmd
}

func newTaskLogListCmd() *cobra.Command {
	var task, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List task logs",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.GetTaskLogs(
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(task))
			if err != nil {
				return err
			}
			format := getOutputFormat()
			if format == "json" || format == "yaml" {
				var data any
				if err := json.Unmarshal(result, &data); err != nil {
					return fmt.Errorf("failed to parse response: %w", err)
				}
				utils.PrintOutput(format, data)
			} else {
				fmt.Println(string(result))
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&task, "task", "", "Task ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}
