package api

import (
	"net/http"
	"testing"
)

func TestCreateParameterTuningJobSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":         1,
			"data":       1,
			"split":      1,
			"evaluation": 1,
		})
	})
	defer server.Close()

	result, err := client.CreateParameterTuningJob(1, 1, 1, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Data != 1 {
		t.Errorf("expected data 1, got %d", result.Data)
	}
	if result.Split != 1 {
		t.Errorf("expected split 1, got %d", result.Split)
	}
	if result.Evaluation != 1 {
		t.Errorf("expected evaluation 1, got %d", result.Evaluation)
	}
}

func TestDeleteParameterTuningJobSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteParameterTuningJob(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetParameterTuningJobsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"count": 1,
			"results": []map[string]any{
				{
					"id":         1,
					"data":       1,
					"split":      1,
					"evaluation": 1,
				},
			},
		})
	})
	defer server.Close()

	result, err := client.GetParameterTuningJobs(nil, nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Count == nil || *result.Count != 1 {
		t.Errorf("expected count 1, got %v", result.Count)
	}
	if result.Results == nil {
		t.Fatal("expected non-nil results")
	}
	if len(*result.Results) != 1 {
		t.Fatalf("expected 1 result, got %d", len(*result.Results))
	}
	job := (*result.Results)[0]
	if job.Id == nil || *job.Id != 1 {
		t.Errorf("expected job id 1, got %v", job.Id)
	}
	if job.Data != 1 {
		t.Errorf("expected job data 1, got %d", job.Data)
	}
	if job.Split != 1 {
		t.Errorf("expected job split 1, got %d", job.Split)
	}
	if job.Evaluation != 1 {
		t.Errorf("expected job evaluation 1, got %d", job.Evaluation)
	}
}
