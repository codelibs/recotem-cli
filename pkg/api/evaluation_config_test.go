package api

import (
	"net/http"
	"testing"

	"recotem.org/cli/recotem/pkg/openapi"
)

func TestCreateEvaluationConfigSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":            1,
			"name":          "eval1",
			"cutoff":        10,
			"target_metric": "ndcg",
		})
	})
	defer server.Close()

	targetMetric := openapi.TargetMetricEnum("ndcg")
	result, err := client.CreateEvaluationConfig(stringPtr("eval1"), intPtr(10), &targetMetric)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name == nil || *result.Name != "eval1" {
		t.Errorf("expected name 'eval1', got %v", result.Name)
	}
	if result.Cutoff == nil || *result.Cutoff != 10 {
		t.Errorf("expected cutoff 10, got %v", result.Cutoff)
	}
	if result.TargetMetric == nil || *result.TargetMetric != openapi.TargetMetricEnum("ndcg") {
		t.Errorf("expected target_metric 'ndcg', got %v", result.TargetMetric)
	}
}

func TestDeleteEvaluationConfigSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteEvaluationConfig(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetEvaluationConfigsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, []map[string]any{
			{
				"id":   1,
				"name": "eval1",
			},
		})
	})
	defer server.Close()

	configs, err := client.GetEvaluationConfigs(nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if configs == nil {
		t.Fatal("expected non-nil configs")
	}
	if len(*configs) != 1 {
		t.Fatalf("expected 1 config, got %d", len(*configs))
	}
	if (*configs)[0].Id == nil || *(*configs)[0].Id != 1 {
		t.Errorf("expected id 1, got %v", (*configs)[0].Id)
	}
	if (*configs)[0].Name == nil || *(*configs)[0].Name != "eval1" {
		t.Errorf("expected name 'eval1', got %v", (*configs)[0].Name)
	}
}

func TestUpdateEvaluationConfigSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":            1,
			"name":          "eval1-updated",
			"cutoff":        20,
			"target_metric": "map",
		})
	})
	defer server.Close()

	targetMetric := openapi.TargetMetricEnum("map")
	result, err := client.UpdateEvaluationConfig(1, stringPtr("eval1-updated"), intPtr(20), &targetMetric)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name == nil || *result.Name != "eval1-updated" {
		t.Errorf("expected name 'eval1-updated', got %v", result.Name)
	}
	if result.Cutoff == nil || *result.Cutoff != 20 {
		t.Errorf("expected cutoff 20, got %v", result.Cutoff)
	}
}
