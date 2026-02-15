package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newAbTestCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ab-test",
		Aliases: []string{"ab"},
		Short:   "Manage A/B tests",
	}

	cmd.AddCommand(
		newAbTestListCmd(),
		newAbTestCreateCmd(),
		newAbTestGetCmd(),
		newAbTestUpdateCmd(),
		newAbTestDeleteCmd(),
		newAbTestStartCmd(),
		newAbTestStopCmd(),
		newAbTestResultsCmd(),
		newAbTestPromoteWinnerCmd(),
	)

	return cmd
}

func newAbTestListCmd() *cobra.Command {
	var project, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List A/B tests",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.GetAbTests(
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
				for _, t := range *result.Results {
					printAbTest(&t)
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

func newAbTestCreateCmd() *cobra.Command {
	var name, project, slots string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create an A/B test",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			projectInt, err := strconv.Atoi(project)
			if err != nil {
				return err
			}
			slotIDs, err := parseIntList(slots)
			if err != nil {
				return fmt.Errorf("invalid slots: %w", err)
			}
			test, err := client.CreateAbTest(name, projectInt, slotIDs)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), test)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "A/B test name")
	cmd.Flags().StringVar(&project, "project", "", "Project ID")
	cmd.Flags().StringVar(&slots, "slots", "", "Deployment slot IDs (comma-separated)")
	_ = cmd.MarkFlagRequired("name")
	_ = cmd.MarkFlagRequired("project")
	_ = cmd.MarkFlagRequired("slots")

	return cmd
}

func newAbTestGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get an A/B test",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			test, err := client.GetAbTest(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), test)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "A/B test ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newAbTestUpdateCmd() *cobra.Command {
	var id, name string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update an A/B test",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			test, err := client.UpdateAbTest(idInt, utils.NilOrString(name), nil)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), test)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "A/B test ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "A/B test name")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newAbTestDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an A/B test",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteAbTest(idInt)
			if err != nil {
				return err
			}
			utils.PrintId(getOutputFormat(), idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "A/B test ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newAbTestStartCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "start",
		Short: "Start an A/B test",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.StartAbTest(idInt)
			if err != nil {
				return err
			}
			fmt.Println("A/B test started.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "A/B test ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newAbTestStopCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "stop",
		Short: "Stop an A/B test",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.StopAbTest(idInt)
			if err != nil {
				return err
			}
			fmt.Println("A/B test stopped.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "A/B test ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newAbTestResultsCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "results",
		Short: "Get A/B test results",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			results, err := client.GetAbTestResults(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), results)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "A/B test ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newAbTestPromoteWinnerCmd() *cobra.Command {
	var id, slotID string

	cmd := &cobra.Command{
		Use:   "promote-winner",
		Short: "Promote A/B test winner",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			slotIDInt, err := strconv.Atoi(slotID)
			if err != nil {
				return err
			}
			result, err := client.PromoteAbTestWinner(idInt, slotIDInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), result)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "A/B test ID")
	cmd.Flags().StringVar(&slotID, "slot-id", "", "Winning slot ID")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("slot-id")

	return cmd
}

func printAbTest(t *openapi.AbTest) {
	id := 0
	if t.Id != nil {
		id = *t.Id
	}
	fmt.Println(id, t.Name, t.Project, t.Status)
}

// parseIntList parses a comma-separated string of integers
func parseIntList(s string) ([]int, error) {
	parts := strings.Split(s, ",")
	result := make([]int, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		v, err := strconv.Atoi(p)
		if err != nil {
			return nil, fmt.Errorf("invalid integer %q: %w", p, err)
		}
		result = append(result, v)
	}
	return result, nil
}
