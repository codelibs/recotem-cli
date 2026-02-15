package api

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetRetrainingRunsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/retraining-run") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"count":    1,
			"next":     nil,
			"previous": nil,
			"results": []map[string]any{
				{
					"id":            1,
					"schedule":      1,
					"status":        "completed",
					"started_at":    nil,
					"completed_at":  nil,
					"error_message": nil,
					"trained_model": nil,
				},
			},
		})
	})
	defer server.Close()

	result, err := client.GetRetrainingRuns(nil, nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Count == nil || *result.Count != 1 {
		t.Errorf("expected count 1, got %v", result.Count)
	}
	if result.Results == nil || len(*result.Results) != 1 {
		t.Fatal("expected 1 result")
	}
	run := (*result.Results)[0]
	if run.Id == nil || *run.Id != 1 {
		t.Errorf("expected id 1, got %v", run.Id)
	}
	if run.Schedule != 1 {
		t.Errorf("expected schedule 1, got %d", run.Schedule)
	}
	if string(run.Status) != "completed" {
		t.Errorf("expected status 'completed', got %s", run.Status)
	}
}

func TestGetRetrainingRunSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/retraining-run") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":            1,
			"schedule":      1,
			"status":        "completed",
			"started_at":    nil,
			"completed_at":  nil,
			"error_message": nil,
			"trained_model": nil,
		})
	})
	defer server.Close()

	result, err := client.GetRetrainingRun(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if run := result; run.Schedule != 1 {
		t.Errorf("expected schedule 1, got %d", run.Schedule)
	}
	if string(result.Status) != "completed" {
		t.Errorf("expected status 'completed', got %s", result.Status)
	}
}
