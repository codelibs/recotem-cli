package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newProjectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Aliases: []string{"p"},
		Short:   "Manage projects",
	}

	cmd.AddCommand(
		newProjectListCmd(),
		newProjectCreateCmd(),
		newProjectDeleteCmd(),
		newProjectSummaryCmd(),
	)

	return cmd
}

func newProjectListCmd() *cobra.Command {
	var id, name string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			projects, err := client.GetProjects(
				utils.NilOrInt(id),
				utils.NilOrString(name))
			if err != nil {
				return err
			}
			for _, x := range *projects {
				printProject(getOutputFormat(), x)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Project ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Project name")

	return cmd
}

func newProjectCreateCmd() *cobra.Command {
	var name, userColumn, itemColumn, timeColumn string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new project",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			project, err := client.CreateProject(name, userColumn, itemColumn,
				utils.NilOrString(timeColumn))
			if err != nil {
				return err
			}
			printProject(getOutputFormat(), *project)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Project name")
	cmd.Flags().StringVarP(&userColumn, "user-column", "u", "", "User column")
	cmd.Flags().StringVarP(&itemColumn, "item-column", "i", "", "Item column")
	cmd.Flags().StringVarP(&timeColumn, "time-column", "t", "", "Time column")
	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("user-column")
	_ = cmd.MarkFlagRequired("item-column")

	return cmd
}

func newProjectDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a project",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteProject(idInt)
			if err != nil {
				return err
			}
			utils.PrintId(getOutputFormat(), idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Project ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newProjectSummaryCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Get project summary",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			summary, err := client.GetProjectSummary(idInt)
			if err != nil {
				return err
			}
			fmt.Println(string(summary))
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Project ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printProject(format string, x openapi.Project) {
	if format == "json" || format == "yaml" {
		m := map[string]any{
			"id":          x.Id,
			"name":        x.Name,
			"user_column": x.UserColumn,
			"item_column": x.ItemColumn,
			"time_column": utils.Atoa(x.TimeColumn),
		}
		utils.PrintOutput(format, m)
	} else {
		fmt.Println(x.Id,
			x.Name,
			x.UserColumn,
			x.ItemColumn,
			utils.Atoa(x.TimeColumn))
	}
}
