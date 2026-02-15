package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newItemMetaDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "item-meta-data",
		Aliases: []string{"imd"},
		Short:   "Manage item meta data",
	}

	cmd.AddCommand(
		newItemMetaDataListCmd(),
		newItemMetaDataUploadCmd(),
		newItemMetaDataDeleteCmd(),
		newItemMetaDataDownloadCmd(),
	)

	return cmd
}

func newItemMetaDataListCmd() *cobra.Command {
	var id, page, pageSize, project string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List item meta data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			tdList, err := client.GetItemMetaData(
				utils.NilOrInt(id),
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(project))
			if err != nil {
				return err
			}
			for _, x := range *tdList.Results {
				printItemMetaData(x)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Item meta data ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")
	cmd.Flags().StringVar(&project, "project", "", "Project ID")

	return cmd
}

func newItemMetaDataUploadCmd() *cobra.Command {
	var project, file string

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload item meta data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			id, err := strconv.Atoi(project)
			if err != nil {
				return err
			}
			itemMetaData, err := client.UploadItemMetaData(id, file)
			if err != nil {
				return err
			}
			printItemMetaData(*itemMetaData)
			return nil
		},
	}

	cmd.Flags().StringVarP(&project, "project", "p", "", "Project ID")
	cmd.Flags().StringVarP(&file, "file", "f", "", "File path")
	_ = cmd.MarkFlagRequired("project")
	_ = cmd.MarkFlagRequired("file")

	return cmd
}

func newItemMetaDataDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete item meta data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteItemMetaData(idInt)
			if err != nil {
				return err
			}
			fmt.Println(idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Item meta data ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newItemMetaDataDownloadCmd() *cobra.Command {
	var id, output string

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download item meta data file",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DownloadItemMetaData(idInt, output)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Item meta data ID")
	cmd.Flags().StringVarP(&output, "output", "O", "", "Output filename")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("output")

	return cmd
}

func printItemMetaData(x openapi.ItemMetaData) {
	fmt.Println(x.Id,
		utils.Atoa(x.Basename),
		x.Filesize,
		utils.FormatTime(x.InsDatetime))
}
