package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newApiKeyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "api-key",
		Aliases: []string{"ak"},
		Short:   "Manage API keys",
	}

	cmd.AddCommand(
		newApiKeyListCmd(),
		newApiKeyCreateCmd(),
		newApiKeyGetCmd(),
		newApiKeyRevokeCmd(),
		newApiKeyDeleteCmd(),
	)

	return cmd
}

func newApiKeyListCmd() *cobra.Command {
	var page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List API keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.GetApiKeys(
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize))
			if err != nil {
				return err
			}
			format := getOutputFormat()
			if format == "json" || format == "yaml" {
				utils.PrintOutput(format, result)
			} else if result.Results != nil {
				for _, k := range *result.Results {
					printApiKey(&k)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newApiKeyCreateCmd() *cobra.Command {
	var name string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new API key",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			key, err := client.CreateApiKey(name)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), key)
			return nil
		},
	}

	cmd.Flags().StringVarP(&name, "name", "n", "", "API key name")
	_ = cmd.MarkFlagRequired("name")

	return cmd
}

func newApiKeyGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get API key details",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			key, err := client.GetApiKey(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), key)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "API key ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newApiKeyRevokeCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "revoke",
		Short: "Revoke an API key",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.RevokeApiKey(idInt)
			if err != nil {
				return err
			}
			fmt.Println("API key revoked.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "API key ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newApiKeyDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete an API key",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteApiKey(idInt)
			if err != nil {
				return err
			}
			utils.PrintId(getOutputFormat(), idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "API key ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printApiKey(k *openapi.ApiKey) {
	prefix := ""
	if k.Prefix != nil {
		prefix = *k.Prefix
	}
	id := 0
	if k.Id != nil {
		id = *k.Id
	}
	fmt.Println(id, k.Name, prefix, k.IsActive)
}
