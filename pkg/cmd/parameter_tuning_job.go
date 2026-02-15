package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"recotem.org/cli/recotem/pkg/openapi"
	"recotem.org/cli/recotem/pkg/utils"
)

func newParameterTuningJobCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "parameter-tuning-job",
		Aliases: []string{"ptj"},
		Short:   "Manage parameter tuning jobs",
	}

	cmd.AddCommand(
		newParameterTuningJobListCmd(),
		newParameterTuningJobCreateCmd(),
		newParameterTuningJobDeleteCmd(),
	)

	return cmd
}

func newParameterTuningJobListCmd() *cobra.Command {
	var data, dataProject, id, page, pageSize string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List parameter tuning jobs",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			ptjList, err := client.GetParameterTuningJobs(
				utils.NilOrInt(data),
				utils.NilOrInt(dataProject),
				utils.NilOrInt(id),
				utils.NilOrInt(page),
				utils.NilOrInt(pageSize))
			if err != nil {
				return err
			}
			for _, x := range *ptjList.Results {
				printParameterTuningJob(x)
			}
			return nil
		},
	}

	cmd.Flags().StringVarP(&data, "data", "d", "", "Data ID")
	cmd.Flags().StringVar(&dataProject, "data-project", "", "Data project ID")
	cmd.Flags().StringVarP(&id, "id", "i", "", "Parameter tuning job ID")
	cmd.Flags().StringVarP(&page, "page", "p", "", "Page")
	cmd.Flags().StringVar(&pageSize, "page-size", "", "Page size")

	return cmd
}

func newParameterTuningJobCreateCmd() *cobra.Command {
	var data, split, evaluation string
	var nTasksParallel, nTrials, memoryBudget, timeoutOverall, timeoutSinglestep, randomSeed string
	var triedAlgorithmJSON, irspackVersion, trainAfterTuning, bestScore, tunedModel, bestConfig string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a parameter tuning job",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			dataID, err := strconv.Atoi(data)
			if err != nil {
				return err
			}
			splitID, err := strconv.Atoi(split)
			if err != nil {
				return err
			}
			evalID, err := strconv.Atoi(evaluation)
			if err != nil {
				return err
			}
			ptj, err := client.CreateParameterTuningJob(
				dataID, splitID, evalID,
				utils.NilOrInt(nTasksParallel),
				utils.NilOrInt(nTrials),
				utils.NilOrInt(memoryBudget),
				utils.NilOrInt(timeoutOverall),
				utils.NilOrInt(timeoutSinglestep),
				utils.NilOrInt(randomSeed),
				utils.NilOrString(triedAlgorithmJSON),
				utils.NilOrString(irspackVersion),
				utils.NilOrBool(trainAfterTuning),
				utils.NilOrFloat32(bestScore),
				utils.NilOrInt(tunedModel),
				utils.NilOrInt(bestConfig))
			if err != nil {
				return err
			}
			printParameterTuningJob(*ptj)
			return nil
		},
	}

	cmd.Flags().StringVarP(&data, "data", "d", "", "Data ID")
	cmd.Flags().StringVarP(&split, "split", "s", "", "Split ID")
	cmd.Flags().StringVarP(&evaluation, "evaluation", "e", "", "Evaluation ID")
	cmd.Flags().StringVar(&nTasksParallel, "n-tasks-parallel", "", "N tasks parallel")
	cmd.Flags().StringVar(&nTrials, "n-trials", "", "N trials")
	cmd.Flags().StringVar(&memoryBudget, "memory-budget", "", "Memory budget")
	cmd.Flags().StringVar(&timeoutOverall, "timeout-overall", "", "Timeout overall")
	cmd.Flags().StringVar(&timeoutSinglestep, "timeout-singlestep", "", "Timeout singlestep")
	cmd.Flags().StringVar(&randomSeed, "random-seed", "", "Random seed")
	cmd.Flags().StringVar(&triedAlgorithmJSON, "tried-algorithm-json", "", "Tried algorithm JSON")
	cmd.Flags().StringVar(&irspackVersion, "irspack-version", "", "irspack version")
	cmd.Flags().StringVar(&trainAfterTuning, "train-after-tuning", "", "Train after tuning")
	cmd.Flags().StringVar(&bestScore, "best-score", "", "Best score")
	cmd.Flags().StringVar(&tunedModel, "tuned-model", "", "Tuned model ID")
	cmd.Flags().StringVar(&bestConfig, "best-config", "", "Best config ID")
	_ = cmd.MarkFlagRequired("data")
	_ = cmd.MarkFlagRequired("split")
	_ = cmd.MarkFlagRequired("evaluation")

	return cmd
}

func newParameterTuningJobDeleteCmd() *cobra.Command {
	var id string

	cmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a parameter tuning job",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := newClientFromCmd(cmd)
			if err != nil {
				return err
			}
			idInt, err := strconv.Atoi(id)
			if err != nil {
				return err
			}
			err = client.DeleteParameterTuningJob(idInt)
			if err != nil {
				return err
			}
			fmt.Println(idInt)
			return nil
		},
	}

	cmd.Flags().StringVarP(&id, "id", "i", "", "Parameter tuning job ID")
	_ = cmd.MarkFlagRequired("id")

	return cmd
}

func printParameterTuningJob(x openapi.ParameterTuningJob) {
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
			status,
			utils.Itoa(x.TunedModel))
	} else {
		fmt.Println(x.Id)
	}
}
