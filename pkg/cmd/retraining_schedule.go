package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newRetrainingScheduleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "retraining-schedule",
		Aliases: []string{"rs"},
		Short:   "Manage retraining schedules",
	}

	cmd.AddCommand(
		newRetrainingScheduleListCmd(),
		newRetrainingScheduleCreateCmd(),
		newRetrainingScheduleGetCmd(),
		newRetrainingScheduleUpdateCmd(),
		newRetrainingScheduleDeleteCmd(),
		newRetrainingScheduleTriggerCmd(),
	)

	return cmd
}

func newRetrainingScheduleListCmd() *cobra.Command {
	var deploymentSlot, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List retraining schedules",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.GetRetrainingSchedules(
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(deploymentSlot))
			if err != nil {
				return err
			}
			format := getOutputFormat()
			if format == "json" || format == "yaml" {
				utils.PrintOutput(format, result)
			} else if result.Results != nil {
				for _, s := range *result.Results {
					printRetrainingSchedule(&s)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&deploymentSlot, "deployment-slot", "", "Deployment slot ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newRetrainingScheduleCreateCmd() *cobra.Command {
	var deploymentSlot, cronExpression string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a retraining schedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			dsInt, err := strconv.Atoi(deploymentSlot)
			if err != nil {
				return err
			}
			schedule, err := client.CreateRetrainingSchedule(dsInt, cronExpression, true)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), schedule)
			return nil
		},
	}

	cmd.Flags().StringVar(&deploymentSlot, "deployment-slot", "", "Deployment slot ID")
	cmd.Flags().StringVar(&cronExpression, "cron-expression", "", "Cron expression")
	_ = cmd.MarkFlagRequired("deployment-slot")
	_ = cmd.MarkFlagRequired("cron-expression")

	return cmd
}

func newRetrainingScheduleGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a retraining schedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			schedule, err := client.GetRetrainingSchedule(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), schedule)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Retraining schedule ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newRetrainingScheduleUpdateCmd() *cobra.Command {
	var id, cronExpression, isActive string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a retraining schedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			schedule, err := client.UpdateRetrainingSchedule(
				idInt,
				utils.NilOrString(cronExpression),
				utils.NilOrBool(isActive))
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), schedule)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Retraining schedule ID")
	cmd.Flags().StringVar(&cronExpression, "cron-expression", "", "Cron expression")
	cmd.Flags().StringVar(&isActive, "is-active", "", "Is active (true/false)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newRetrainingScheduleDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a retraining schedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteRetrainingSchedule(idInt)
			if err != nil {
				return err
			}
			utils.PrintId(getOutputFormat(), idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Retraining schedule ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newRetrainingScheduleTriggerCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "trigger",
		Short: "Trigger a retraining schedule",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.TriggerRetrainingSchedule(idInt)
			if err != nil {
				return err
			}
			fmt.Println("Retraining triggered.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Retraining schedule ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printRetrainingSchedule(s *openapi.RetrainingSchedule) {
	id := 0
	if s.Id != nil {
		id = *s.Id
	}
	fmt.Println(id, s.DeploymentSlot, s.CronExpression, s.IsActive)
}
