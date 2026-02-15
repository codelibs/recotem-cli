package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newTrainingDataCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "training-data",
		Aliases: []string{"td"},
		Short:   "Manage training data",
	}

	cmd.AddCommand(
		newTrainingDataListCmd(),
		newTrainingDataUploadCmd(),
		newTrainingDataDeleteCmd(),
		newTrainingDataDownloadCmd(),
		newTrainingDataPreviewCmd(),
	)

	return cmd
}

func newTrainingDataListCmd() *cobra.Command {
	var id, page, pageSize, project string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List training data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			tdList, err := client.GetTrainingData(
				utils.NilOrInt(id),
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize),
				utils.NilOrInt(project))
			if err != nil {
				return err
			}
			for _, x := range *tdList.Results {
				printTrainingData(x)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Training data ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")
	cmd.Flags().StringVar(&project, "project", "", "Project ID")

	return cmd
}

func newTrainingDataUploadCmd() *cobra.Command {
	var project, file string

	cmd := &cobra.Command{
		Use:   "upload",
		Short: "Upload training data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			id, err := strconv.Atoi(project)
			if err != nil {
				return err
			}
			trainingData, err := client.UploadTrainingData(id, file)
			if err != nil {
				return err
			}
			printTrainingData(*trainingData)
			return nil
		},
	}

	cmd.Flags().StringVarP(&project, "project", "p", "", "Project ID")
	cmd.Flags().StringVarP(&file, "file", "f", "", "File path")
	_ = cmd.MarkFlagRequired("project")
	_ = cmd.MarkFlagRequired("file")

	return cmd
}

func newTrainingDataDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete training data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteTrainingData(idInt)
			if err != nil {
				return err
			}
			fmt.Println(idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Training data ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTrainingDataDownloadCmd() *cobra.Command {
	var id, output string

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download training data file",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DownloadTrainingData(idInt, output)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Training data ID")
	cmd.Flags().StringVarP(&output, "output", "O", "", "Output filename")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("output")

	return cmd
}

func newTrainingDataPreviewCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "preview",
		Short: "Preview training data",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			preview, err := client.PreviewTrainingData(idInt)
			if err != nil {
				return err
			}
			fmt.Println(string(preview))
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Training data ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printTrainingData(x openapi.TrainingData) {
	fmt.Println(x.Id,
		x.Project,
		utils.Atoa(x.Basename),
		x.Filesize,
		utils.FormatTime(x.InsDatetime))
}
