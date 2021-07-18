package api

import (
	"fmt"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) CreateParameterTuningJob(data int, split int, evaluation int, nTasksParallel *int,
	nTrials *int, memoryBudget *int, timeoutOverall *int, timeoutSinglestep *int, randomeSeed *int,
	triedAlgorithmsJson *string, irspackVersion *string, trainAfterTuning *bool, bestScore *float32,
	tunedModel *int, bestConfig *int) (*openapi.ParameterTuningJob, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	req := openapi.ParameterTuningJobCreateJSONRequestBody{
		Data:                data,
		Split:               split,
		Evaluation:          evaluation,
		NTasksParallel:      nTasksParallel,
		NTrials:             nTrials,
		MemoryBudget:        memoryBudget,
		TimeoutOverall:      timeoutOverall,
		TimeoutSinglestep:   timeoutSinglestep,
		RandomSeed:          randomeSeed,
		TriedAlgorithmsJson: triedAlgorithmsJson,
		IrspackVersion:      irspackVersion,
		TrainAfterTuning:    trainAfterTuning,
		BestScore:           bestScore,
		TunedModel:          tunedModel,
		BestConfig:          bestConfig,
	}
	resp, err := client.ParameterTuningJobCreateWithResponse(c.Context, req)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}

func (c Client) GetParameterTuningJobs(data *int, dataProject *int, id *int, page *int,
	pageSize *int) (*openapi.PaginatedParameterTuningJobList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.ParameterTuningJobListParams
	if data != nil || dataProject != nil || id != nil || page != nil || pageSize != nil {
		req = openapi.ParameterTuningJobListParams{
			Data:        data,
			DataProject: dataProject,
			Id:          id,
			Page:        page,
			PageSize:    pageSize,
		}
	}
	resp, err := client.ParameterTuningJobListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}
