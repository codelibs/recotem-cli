package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newSplitConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "split-config",
		Aliases: []string{"sc"},
		Short:   "Manage split configs",
	}

	cmd.AddCommand(
		newSplitConfigListCmd(),
		newSplitConfigCreateCmd(),
		newSplitConfigDeleteCmd(),
		newSplitConfigUpdateCmd(),
	)

	return cmd
}

func newSplitConfigListCmd() *cobra.Command {
	var id, name, unnamed string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List split configs",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			configs, err := client.GetSplitConfigs(
				utils.NilOrInt(id),
				utils.NilOrString(name),
				utils.NilOrBool(unnamed))
			if err != nil {
				return err
			}
			for _, x := range *configs {
				printSplitConfig(getOutputFormat(), x)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Split config ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&unnamed, "unnamed", "u", "", "Unnamed")

	return cmd
}

func newSplitConfigCreateCmd() *cobra.Command {
	var name, scheme, heldoutRatio, nHeldout, testUserRatio, nTestUsers, randomSeed string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a split config",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			sc, err := client.CreateSplitConfig(
				utils.NilOrString(name),
				utils.NilOrScheme(scheme),
				utils.NilOrFloat32(heldoutRatio),
				utils.NilOrInt(nHeldout),
				utils.NilOrFloat32(testUserRatio),
				utils.NilOrInt(nTestUsers),
				utils.NilOrInt(randomSeed))
			if err != nil {
				return err
			}
			printSplitConfig(getOutputFormat(), *sc)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&scheme, "scheme", "s", "", "Scheme (RG|TG|TU)")
	cmd.Flags().StringVar(&heldoutRatio, "heldout-ratio", "", "Heldout ratio")
	cmd.Flags().StringVar(&nHeldout, "n-heldout", "", "Number of heldout")
	cmd.Flags().StringVar(&testUserRatio, "test-user-ratio", "", "Test user ratio")
	cmd.Flags().StringVar(&nTestUsers, "n-test-users", "", "Number of test users")
	cmd.Flags().StringVar(&randomSeed, "random-seed", "", "Random seed")

	return cmd
}

func newSplitConfigDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a split config",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteSplitConfig(idInt)
			if err != nil {
				return err
			}
			utils.PrintId(getOutputFormat(), idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Split config ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newSplitConfigUpdateCmd() *cobra.Command {
	var id, name, scheme, heldoutRatio, nHeldout, testUserRatio, nTestUsers, randomSeed string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a split config",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			sc, err := client.UpdateSplitConfig(idInt,
				utils.NilOrString(name),
				utils.NilOrScheme(scheme),
				utils.NilOrFloat32(heldoutRatio),
				utils.NilOrInt(nHeldout),
				utils.NilOrFloat32(testUserRatio),
				utils.NilOrInt(nTestUsers),
				utils.NilOrInt(randomSeed))
			if err != nil {
				return err
			}
			printSplitConfig(getOutputFormat(), *sc)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Split config ID")
	cmd.Flags().StringVarP(&name, "name", "n", "", "Name")
	cmd.Flags().StringVarP(&scheme, "scheme", "s", "", "Scheme (RG|TG|TU)")
	cmd.Flags().StringVar(&heldoutRatio, "heldout-ratio", "", "Heldout ratio")
	cmd.Flags().StringVar(&nHeldout, "n-heldout", "", "Number of heldout")
	cmd.Flags().StringVar(&testUserRatio, "test-user-ratio", "", "Test user ratio")
	cmd.Flags().StringVar(&nTestUsers, "n-test-users", "", "Number of test users")
	cmd.Flags().StringVar(&randomSeed, "random-seed", "", "Random seed")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printSplitConfig(format string, x openapi.SplitConfig) {
	if format == "json" || format == "yaml" {
		m := map[string]any{
			"id":              x.Id,
			"heldout_ratio":   utils.Ftoa(x.HeldoutRatio),
			"test_user_ratio": utils.Ftoa(x.TestUserRatio),
			"random_seed":     utils.Itoa(x.RandomSeed),
		}
		utils.PrintOutput(format, m)
	} else {
		fmt.Println(x.Id,
			utils.Ftoa(x.HeldoutRatio),
			utils.Ftoa(x.TestUserRatio),
			utils.Itoa(x.RandomSeed))
	}
}
