package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newConversionEventCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "conversion-event",
		Aliases: []string{"ce"},
		Short:   "Manage conversion events",
	}

	cmd.AddCommand(
		newConversionEventListCmd(),
		newConversionEventCreateCmd(),
		newConversionEventBatchCreateCmd(),
		newConversionEventGetCmd(),
	)

	return cmd
}

func newConversionEventListCmd() *cobra.Command {
	var abTest, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List conversion events",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.GetConversionEvents(
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(abTest))
			if err != nil {
				return err
			}
			format := getOutputFormat()
			if format == "json" || format == "yaml" {
				utils.PrintOutput(format, result)
			} else if result.Results != nil {
				for _, e := range *result.Results {
					printConversionEvent(&e)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&abTest, "ab-test", "", "A/B test ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newConversionEventCreateCmd() *cobra.Command {
	var abTest, slot, userID, itemID, eventType string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a conversion event",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			abTestInt, err := strconv.Atoi(abTest)
			if err != nil {
				return err
			}
			slotInt, err := strconv.Atoi(slot)
			if err != nil {
				return err
			}
			event, err := client.CreateConversionEvent(
				abTestInt,
				userID,
				utils.NilOrString(itemID),
				eventType,
				slotInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), event)
			return nil
		},
	}

	cmd.Flags().StringVar(&abTest, "ab-test", "", "A/B test ID")
	cmd.Flags().StringVar(&slot, "slot", "", "Deployment slot ID")
	cmd.Flags().StringVar(&userID, "user-id", "", "User ID")
	cmd.Flags().StringVar(&itemID, "item-id", "", "Item ID")
	cmd.Flags().StringVar(&eventType, "event-type", "", "Event type")
	_ = cmd.MarkFlagRequired("ab-test")
	_ = cmd.MarkFlagRequired("slot")
	_ = cmd.MarkFlagRequired("user-id")
	_ = cmd.MarkFlagRequired("event-type")

	return cmd
}

func newConversionEventBatchCreateCmd() *cobra.Command {
	var file string

	cmd := &cobra.Command{
		Use:   "batch-create",
		Short: "Batch create conversion events from JSON file",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			data, err := os.ReadFile(file)
			if err != nil {
				return fmt.Errorf("failed to read file: %w", err)
			}
			var events []openapi.ConversionEventCreate
			if unmarshalErr := json.Unmarshal(data, &events); unmarshalErr != nil {
				return fmt.Errorf("failed to parse JSON: %w", unmarshalErr)
			}
			result, err := client.BatchCreateConversionEvents(events)
			if err != nil {
				return err
			}
			fmt.Println(string(result))
			return nil
		},
	}

	cmd.Flags().StringVarP(&file, "file", "f", "", "JSON file path")
	_ = cmd.MarkFlagRequired("file")

	return cmd
}

func newConversionEventGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a conversion event",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			event, err := client.GetConversionEvent(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), event)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Conversion event ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printConversionEvent(e *openapi.ConversionEvent) {
	id := 0
	if e.Id != nil {
		id = *e.Id
	}
	fmt.Println(id, e.AbTest, e.Slot, e.UserId, e.EventType)
}
