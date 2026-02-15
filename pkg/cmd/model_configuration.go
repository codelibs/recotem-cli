package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newModelConfigurationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "model-configuration",
		Aliases: []string{"mc"},
		Short:   "Manage model configurations",
	}

	cmd.AddCommand(
		newModelConfigurationListCmd(),
		newModelConfigurationCreateCmd(),
		newModelConfigurationDeleteCmd(),
		newModelConfigurationUpdateCmd(),
	)

	return cmd
}

func newModelConfigurationListCmd() *cobra.Command {
	var id, page, pageSize, project string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List model configurations",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			modelConfigs, err := client.GetModelConfigurations(
				utils.NilOrInt(id),
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(project))
			if err != nil {
				return err
			}
			for _, x := range *modelConfigs.Results {
				printModelConfiguration(x)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Model configuration ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")
	cmd.Flags().StringVar(&project, "project", "", "Project ID")

	return cmd
}

func newModelConfigurationCreateCmd() *cobra.Command {
	var name, project, recommenderClassName, parametersJSON string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a model configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			projectID, err := strconv.Atoi(project)
			if err != nil {
				return err
			}
			mc, err := client.CreateModelConfiguration(
				utils.NilOrString(name), projectID,
				recommenderClassName, parametersJSON)
			if err != nil {
				return err
			}
			printModelConfiguration(*mc)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&project, "project", "p", "", "Project ID")
	cmd.Flags().StringVar(&recommenderClassName, "recommender-class-name", "", "Recommender class name")
	cmd.Flags().StringVar(&parametersJSON, "parameters-json", "", "Parameters JSON")
	_ = cmd.MarkFlagRequired("project")
	_ = cmd.MarkFlagRequired("recommender-class-name")
	_ = cmd.MarkFlagRequired("parameters-json")

	return cmd
}

func newModelConfigurationDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a model configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteModelConfiguration(idInt)
			if err != nil {
				return err
			}
			fmt.Println(idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Model configuration ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newModelConfigurationUpdateCmd() *cobra.Command {
	var id, name, recommenderClassName, parametersJSON string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a model configuration",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			mc, err := client.UpdateModelConfiguration(idInt,
				utils.NilOrString(name),
				utils.NilOrString(parametersJSON))
			if err != nil {
				return err
			}
			printModelConfiguration(*mc)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Model configuration ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVar(&recommenderClassName, "recommender-class-name", "", "Recommender class name")
	cmd.Flags().StringVar(&parametersJSON, "parameters-json", "", "Parameters JSON")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printModelConfiguration(x openapi.ModelConfiguration) {
	fmt.Println(x.Id,
		x.Project,
		x.RecommenderClassName,
		x.TuningJob,
		utils.FormatName(utils.Atoa(x.Name)))
}
