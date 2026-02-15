package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newRetrainingRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "retraining-run",
		Aliases: []string{"rr"},
		Short:   "Manage retraining runs",
	}

	cmd.AddCommand(
		newRetrainingRunListCmd(),
		newRetrainingRunGetCmd(),
	)

	return cmd
}

func newRetrainingRunListCmd() *cobra.Command {
	var schedule, status, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List retraining runs",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			var statusParam *openapi.RetrainingRunListParamsStatus
			if status != "" {
				s := openapi.RetrainingRunListParamsStatus(status)
				statusParam = &s
			}
			result, err := client.GetRetrainingRuns(
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(schedule),
				statusParam)
			if err != nil {
				return err
			}
			format := getOutputFormat()
			if format == "json" || format == "yaml" {
				utils.PrintOutput(format, result)
			} else if result.Results != nil {
				for _, r := range *result.Results {
					printRetrainingRun(&r)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&schedule, "schedule", "", "Schedule ID")
	cmd.Flags().StringVar(&status, "status", "", "Status filter")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newRetrainingRunGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a retraining run",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			run, err := client.GetRetrainingRun(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), run)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Retraining run ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printRetrainingRun(r *openapi.RetrainingRun) {
	id := 0
	if r.Id != nil {
		id = *r.Id
	}
	fmt.Println(id, r.Schedule, r.Status, r.StartedAt, r.CompletedAt)
}
