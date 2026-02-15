package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newUserCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "user",
		Aliases: []string{"u"},
		Short:   "Manage users",
	}

	cmd.AddCommand(
		newUserListCmd(),
		newUserCreateCmd(),
		newUserGetCmd(),
		newUserUpdateCmd(),
		newUserDeactivateCmd(),
		newUserActivateCmd(),
		newUserResetPasswordCmd(),
	)

	return cmd
}

func newUserListCmd() *cobra.Command {
	var page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List users",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			result, err := client.GetUsers(
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize))
			if err != nil {
				return err
			}
			format := getOutputFormat()
			if format == "json" || format == "yaml" {
				utils.PrintOutput(format, result)
			} else if result.Results != nil {
				for _, u := range *result.Results {
					printUser(&u)
				}
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newUserCreateCmd() *cobra.Command {
	var username, email, password string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			user, err := client.CreateUser(username, password, email)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), user)
			return nil
		},
	}

	cmd.Flags().StringVar(&username, "username", "", "Username")
	cmd.Flags().StringVar(&email, "email", "", "Email")
	cmd.Flags().StringVar(&password, "password", "", "Password")
	_ = cmd.MarkFlagRequired("username")
	_ = cmd.MarkFlagRequired("email")
	_ = cmd.MarkFlagRequired("password")

	return cmd
}

func newUserGetCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			user, err := client.GetUser(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), user)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "User ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newUserUpdateCmd() *cobra.Command {
	var id, email, isActive string

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			user, err := client.UpdateUser(
				idInt,
				utils.NilOrString(email),
				utils.NilOrBool(isActive))
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), user)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "User ID")
	cmd.Flags().StringVar(&email, "email", "", "Email")
	cmd.Flags().StringVar(&isActive, "is-active", "", "Is active (true/false)")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newUserDeactivateCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "deactivate",
		Short: "Deactivate a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeactivateUser(idInt)
			if err != nil {
				return err
			}
			fmt.Println("User deactivated.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "User ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newUserActivateCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "activate",
		Short: "Activate a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.ActivateUser(idInt)
			if err != nil {
				return err
			}
			fmt.Println("User activated.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "User ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newUserResetPasswordCmd() *cobra.Command {
	var id, newPassword string

	cmd := &cobra.Command{
		Use:   "reset-password",
		Short: "Reset a user's password",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.ResetUserPassword(idInt, newPassword)
			if err != nil {
				return err
			}
			fmt.Println("Password reset.")
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "User ID")
	cmd.Flags().StringVar(&newPassword, "new-password", "", "New password")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("new-password")

	return cmd
}

func printUser(u *openapi.User) {
	id := 0
	if u.Id != nil {
		id = *u.Id
	}
	fmt.Println(id, u.Username, u.Email, u.IsActive, u.IsStaff)
}
