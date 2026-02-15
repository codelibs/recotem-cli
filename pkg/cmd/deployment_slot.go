package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newDeploymentSlotCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "deployment-slot",
		Aliases: []string{"ds"},
		Short:   "Manage deployment slots",
	}

	cmd.AddCommand(
		newDeploymentSlotListCmd(),
		newDeploymentSlotCreateCmd(),
		newDeploymentSlotGetCmd(),
		newDeploymentSlotUpdateCmd(),
		newDeploymentSlotDeleteCmd(),
	)

	return cmd
}

func newDeploymentSlotListCmd() *cobra.Command {
	var project, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List deployment slots",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.GetDeploymentSlots(
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(project))
			if err != nil {
				return err
			}
			format := getOutputFormat()
			if format == "json" || format == "yaml" {
				utils.PrintOutput(format, result)
			} else if result.Results != nil {
				for _, s := range *result.Results {
					printDeploymentSlot(&s)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&project, "project", "", "Project ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newDeploymentSlotCreateCmd() *cobra.Command {
	var name, project, trainedModel string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a deployment slot",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			projectInt, err := strconv.Atoi(project)
			if err != nil {
				return err
			}
			slot, err := client.CreateDeploymentSlot(
				name,
				projectInt,
				utils.NilOrInt(trainedModel),
				true)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), slot)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Slot name")
	cmd.Flags().StringVar(&project, "project", "", "Project ID")
	cmd.Flags().StringVar(&trainedModel, "trained-model", "", "Trained model ID")
	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("project")

	return cmd
}

func newDeploymentSlotGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a deployment slot",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			slot, err := client.GetDeploymentSlot(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), slot)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Deployment slot ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newDeploymentSlotUpdateCmd() *cobra.Command {
	var id, name, trainedModel string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a deployment slot",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			slot, err := client.UpdateDeploymentSlot(
				idInt,
				utils.NilOrString(name),
				utils.NilOrInt(trainedModel))
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), slot)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Deployment slot ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Slot name")
	cmd.Flags().StringVar(&trainedModel, "trained-model", "", "Trained model ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newDeploymentSlotDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a deployment slot",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteDeploymentSlot(idInt)
			if err != nil {
				return err
			}
			utils.PrintId(getOutputFormat(), idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Deployment slot ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printDeploymentSlot(s *openapi.DeploymentSlot) {
	id := 0
	if s.Id != nil {
		id = *s.Id
	}
	tm := "<none>"
	if s.TrainedModel != nil {
		tm = strconv.Itoa(*s.TrainedModel)
	}
	fmt.Println(id, s.Name, s.Project, tm, s.IsActive)
}
