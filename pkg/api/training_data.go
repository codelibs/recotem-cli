package api

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"

	"recotem.org/cli/recotem/pkg/openapi"
)

func (c Client) UploadTrainingData(projectId int, uploadPath string) (*openapi.TrainingData, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	f, err := os.Open(uploadPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	filename := filepath.Base(uploadPath)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	fileWriter, err := writer.CreateFormFile("file", filename)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(fileWriter, f)
	if err != nil {
		return nil, err
	}

	err = writer.WriteField("project", strconv.Itoa(projectId))
	if err != nil {
		return nil, err
	}

	contentType := writer.FormDataContentType()

	resp, err := client.TrainingDataCreateWithBodyWithResponse(c.Context, contentType, body)
	if err != nil {
		return nil, err
	}

	if resp.JSON201 != nil {
		return resp.JSON201, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DeleteTrainingData(id int) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.TrainingDataDestroyWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return nil
	}

	return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) GetTrainingData(id *int, page *int, pageSize *int, project *int) (*openapi.PaginatedTrainingDataList, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	var req openapi.TrainingDataListParams
	if id != nil || page != nil || pageSize != nil || project != nil {
		req = openapi.TrainingDataListParams{
			Id:       id,
			Page:     page,
			PageSize: pageSize,
			Project:  project,
		}
	}
	resp, err := client.TrainingDataListWithResponse(c.Context, &req)
	if err != nil {
		return nil, err
	}

	if resp.JSON200 != nil {
		return resp.JSON200, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}

func (c Client) DownloadTrainingData(id int, output string) error {
	client, err := c.newApiClient()
	if err != nil {
		return err
	}

	resp, err := client.TrainingDataDownloadWithResponse(c.Context, id)
	if err != nil {
		return err
	}

	if resp.StatusCode() != 200 {
		return fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
	}

	return os.WriteFile(output, resp.Body, 0600)
}

func (c Client) PreviewTrainingData(id int) ([]byte, error) {
	client, err := c.newApiClient()
	if err != nil {
		return nil, err
	}

	resp, err := client.TrainingDataPreviewWithResponse(c.Context, id, nil)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() >= 200 && resp.StatusCode() < 300 {
		return resp.Body, nil
	}

	return nil, fmt.Errorf("%s: %s", resp.Status(), string(resp.Body))
}
