package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newEvaluationConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "evaluation-config",
		Aliases: []string{"ec"},
		Short:   "Manage evaluation configs",
	}

	cmd.AddCommand(
		newEvaluationConfigListCmd(),
		newEvaluationConfigCreateCmd(),
		newEvaluationConfigDeleteCmd(),
		newEvaluationConfigUpdateCmd(),
	)

	return cmd
}

func newEvaluationConfigListCmd() *cobra.Command {
	var id, name, unnamed string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List evaluation configs",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			configs, err := client.GetEvaluationConfigs(
				utils.NilOrInt(id),
				utils.NilOrString(name),
				utils.NilOrBool(unnamed))
			if err != nil {
				return err
			}
			for _, x := range *configs {
				printEvaluationConfig(getOutputFormat(), x)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Evaluation config ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&unnamed, "unnamed", "u", "", "Unnamed")

	return cmd
}

func newEvaluationConfigCreateCmd() *cobra.Command {
	var name, cutoff, targetMetric string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an evaluation config",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			ec, err := client.CreateEvaluationConfig(
				utils.NilOrString(name),
				utils.NilOrInt(cutoff),
				utils.NilOrTargetMetric(targetMetric))
			if err != nil {
				return err
			}
			printEvaluationConfig(getOutputFormat(), *ec)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&cutoff, "cutoff", "c", "", "Cutoff")
	cmd.Flags().StringVar(&targetMetric, "target-metric", "", "Target metric (ndcg|map|recall|hit)")

	return cmd
}

func newEvaluationConfigDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an evaluation config",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteEvaluationConfig(idInt)
			if err != nil {
				return err
			}
			utils.PrintId(getOutputFormat(), idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Evaluation config ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newEvaluationConfigUpdateCmd() *cobra.Command {
	var id, name, cutoff, targetMetric string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an evaluation config",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			ec, err := client.UpdateEvaluationConfig(idInt,
				utils.NilOrString(name),
				utils.NilOrInt(cutoff),
				utils.NilOrTargetMetric(targetMetric))
			if err != nil {
				return err
			}
			printEvaluationConfig(getOutputFormat(), *ec)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Evaluation config ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&cutoff, "cutoff", "c", "", "Cutoff")
	cmd.Flags().StringVar(&targetMetric, "target-metric", "", "Target metric (ndcg|map|recall|hit)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printEvaluationConfig(format string, x openapi.EvaluationConfig) {
	var targetMetric string
	if x.TargetMetric != nil {
		targetMetric = string(*x.TargetMetric)
	} else {
		targetMetric = utils.NoValue
	}
	if format == "json" || format == "yaml" {
		m := map[string]any{
			"id":            x.Id,
			"cutoff":        utils.Itoa(x.Cutoff),
			"target_metric": targetMetric,
			"name":          utils.Atoa(x.Name),
		}
		utils.PrintOutput(format, m)
	} else {
		fmt.Println(
			x.Id,
			utils.Itoa(x.Cutoff),
			targetMetric,
			utils.FormatName(utils.Atoa(x.Name)))
	}
}
