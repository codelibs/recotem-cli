package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newTrainedModelCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "trained-model",
		Aliases: []string{"tm"},
		Short:   "Manage trained models",
	}

	cmd.AddCommand(
		newTrainedModelListCmd(),
		newTrainedModelCreateCmd(),
		newTrainedModelDeleteCmd(),
		newTrainedModelDownloadCmd(),
		newTrainedModelRecommendCmd(),
		newTrainedModelSampleRecommendCmd(),
		newTrainedModelRecommendProfileCmd(),
	)

	return cmd
}

func newTrainedModelListCmd() *cobra.Command {
	var dataLoc, dataLocProject, id, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List trained models",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			tmList, err := client.GetTrainedModels(
				utils.NilOrInt(dataLoc),
				utils.NilOrInt(dataLocProject),
				utils.NilOrInt(id),
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize))
			if err != nil {
				return err
			}
			for _, x := range *tmList.Results {
				printTrainedModel(x)
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&dataLoc, "data-loc", "", "Data loc")
	cmd.Flags().StringVar(&dataLocProject, "data-loc-project", "", "Data loc project")
	cmd.Flags().StringVarP(&id, "id", "i", "", "Trained model ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newTrainedModelCreateCmd() *cobra.Command {
	var configuration, dataLoc, file, irspackVersion string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a trained model",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			configID, err := strconv.Atoi(configuration)
			if err != nil {
				return err
			}
			dataLocID, err := strconv.Atoi(dataLoc)
			if err != nil {
				return err
			}
			trainedModel, err := client.CreateTrainedModel(
				configID, dataLocID,
				utils.NilOrString(file),
				utils.NilOrString(irspackVersion))
			if err != nil {
				return err
			}
			printTrainedModel(*trainedModel)
			return nil
		},
	}

	cmd.Flags().StringVarP(&configuration, "configuration", "c", "", "Configuration ID")
	cmd.Flags().StringVar(&dataLoc, "data-loc", "", "Data loc ID")
	cmd.Flags().StringVarP(&file, "file", "f", "", "File")
	cmd.Flags().StringVar(&irspackVersion, "irspack-version", "", "irspack version")
	_ = cmd.MarkFlagRequired("configuration")
	_ = cmd.MarkFlagRequired("data-loc")

	return cmd
}

func newTrainedModelDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a trained model",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteTrainedModel(idInt)
			if err != nil {
				return err
			}
			fmt.Println(idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Trained model ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTrainedModelDownloadCmd() *cobra.Command {
	var id, output string

	cmd := &cobra.Command{
		Use:   "download",
		Short: "Download trained model file",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DownloadTrainedModel(idInt, output)
			if err != nil {
				return err
			}
			fmt.Println(output)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Trained model ID")
	cmd.Flags().StringVarP(&output, "output", "O", "", "Output filename")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("output")

	return cmd
}

func newTrainedModelRecommendCmd() *cobra.Command {
	var id, userID string
	var nItems int

	cmd := &cobra.Command{
		Use:   "recommend",
		Short: "Get recommendations for a user",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			result, err := client.Recommend(idInt, userID, nItems)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), result)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Trained model ID")
	cmd.Flags().StringVar(&userID, "user-id", "", "User ID")
	cmd.Flags().IntVarP(&nItems, "n-items", "n", 10, "Number of items to recommend")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("user-id")

	return cmd
}

func newTrainedModelSampleRecommendCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "sample-recommend",
		Short: "Get sample recommendations",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			result, err := client.SampleRecommend(idInt)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), result)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Trained model ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func newTrainedModelRecommendProfileCmd() *cobra.Command {
	var id string
	var itemIDs []string
	var nItems int

	cmd := &cobra.Command{
		Use:   "recommend-profile",
		Short: "Get recommendations based on item profile",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			result, err := client.RecommendProfile(idInt, itemIDs, nItems)
			if err != nil {
				return err
			}
			utils.PrintOutput(getOutputFormat(), result)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Trained model ID")
	cmd.Flags().StringSliceVar(&itemIDs, "item-ids", nil, "Item IDs (comma-separated)")
	cmd.Flags().IntVarP(&nItems, "n-items", "n", 10, "Number of items to recommend")
	_ = cmd.MarkFlagRequired("id")
	_ = cmd.MarkFlagRequired("item-ids")

	return cmd
}

func printTrainedModel(x openapi.TrainedModel) {
	if x.TaskLinks != nil && len(*x.TaskLinks) > 0 {
		task := (*x.TaskLinks)[len(*x.TaskLinks)-1].Task
		var status string
		if task.Status != nil {
			status = string(*task.Status)
		} else {
			status = utils.NoValue
		}
		fmt.Println(x.Id,
			utils.FormatTime(x.InsDatetime),
			status)
	} else {
		fmt.Println(x.Id)
	}
}
