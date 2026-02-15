package api

import (
	"net/http"
	"testing"
)

func TestCreateModelConfigurationSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":                     1,
			"name":                   "test-config",
			"project":                1,
			"recommender_class_name": "P3alphaRecommender",
			"parameters_json":        "{}",
		})
	})
	defer server.Close()

	name := "test-config"
	result, err := client.CreateModelConfiguration(&name, 1, "P3alphaRecommender", "{}")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.RecommenderClassName != "P3alphaRecommender" {
		t.Errorf("expected recommender_class_name P3alphaRecommender, got %s", result.RecommenderClassName)
	}
	if result.Project != 1 {
		t.Errorf("expected project 1, got %d", result.Project)
	}
}

func TestDeleteModelConfigurationSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteModelConfiguration(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetModelConfigurationsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"count": 1,
			"results": []map[string]any{
				{
					"id":                     1,
					"name":                   "test-config",
					"project":                1,
					"recommender_class_name": "P3alphaRecommender",
					"parameters_json":        "{}",
				},
			},
		})
	})
	defer server.Close()

	result, err := client.GetModelConfigurations(nil, nil, nil, intPtr(1))
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
	first := (*result.Results)[0]
	if first.Id == nil || *first.Id != 1 {
		t.Errorf("expected id 1, got %v", first.Id)
	}
	if first.RecommenderClassName != "P3alphaRecommender" {
		t.Errorf("expected recommender_class_name P3alphaRecommender, got %s", first.RecommenderClassName)
	}
}

func TestUpdateModelConfigurationSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPatch {
			t.Errorf("expected PATCH method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":                     1,
			"name":                   "updated-config",
			"project":                1,
			"recommender_class_name": "P3alphaRecommender",
			"parameters_json":        "{\"alpha\": 0.5}",
		})
	})
	defer server.Close()

	result, err := client.UpdateModelConfiguration(1, stringPtr("updated-config"), stringPtr("{\"alpha\": 0.5}"))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name == nil || *result.Name != "updated-config" {
		t.Errorf("expected name updated-config, got %v", result.Name)
	}
	if result.ParametersJson != "{\"alpha\": 0.5}" {
		t.Errorf("expected parameters_json {\"alpha\": 0.5}, got %s", result.ParametersJson)
	}
}
