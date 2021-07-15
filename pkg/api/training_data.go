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

	return nil, fmt.Errorf(fmt.Sprintf("%s: %s", resp.Status(), string(resp.Body)))
}
