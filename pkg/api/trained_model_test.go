package api

import (
	"net/http"
	"testing"
)

func TestCreateTrainedModelSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":            1,
			"configuration": 10,
			"data_loc":      20,
		})
	})
	defer server.Close()

	result, err := client.CreateTrainedModel(10, 20, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Configuration != 10 {
		t.Errorf("expected configuration 10, got %d", result.Configuration)
	}
	if result.DataLoc != 20 {
		t.Errorf("expected data_loc 20, got %d", result.DataLoc)
	}
}

func TestDeleteTrainedModelSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteTrainedModel(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetTrainedModelsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"count": 1,
			"results": []map[string]any{
				{
					"id":            1,
					"configuration": 10,
					"data_loc":      20,
				},
			},
		})
	})
	defer server.Close()

	result, err := client.GetTrainedModels(nil, nil, nil, nil, nil)
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
	if first.Configuration != 10 {
		t.Errorf("expected configuration 10, got %d", first.Configuration)
	}
}

func TestRecommendSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"user_id": "u1",
			"recommendations": []map[string]any{
				{"item_id": "i1", "score": 0.9},
				{"item_id": "i2", "score": 0.8},
			},
			"user_profile": []string{},
		})
	})
	defer server.Close()

	result, err := client.Recommend(1, "u1", 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.UserId != "u1" {
		t.Errorf("expected user_id u1, got %s", result.UserId)
	}
	if len(result.Recommendations) != 2 {
		t.Fatalf("expected 2 recommendations, got %d", len(result.Recommendations))
	}
	if result.Recommendations[0].ItemId != "i1" {
		t.Errorf("expected first item_id i1, got %s", result.Recommendations[0].ItemId)
	}
	if result.Recommendations[1].ItemId != "i2" {
		t.Errorf("expected second item_id i2, got %s", result.Recommendations[1].ItemId)
	}
}

func TestSampleRecommendSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"user_id": "u1",
			"recommendations": []map[string]any{
				{"item_id": "i1", "score": 0.9},
				{"item_id": "i2", "score": 0.8},
			},
			"user_profile": []string{},
		})
	})
	defer server.Close()

	result, err := client.SampleRecommend(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.UserId != "u1" {
		t.Errorf("expected user_id u1, got %s", result.UserId)
	}
	if len(result.Recommendations) != 2 {
		t.Fatalf("expected 2 recommendations, got %d", len(result.Recommendations))
	}
}

func TestRecommendProfileSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"user_id": "",
			"recommendations": []map[string]any{
				{"item_id": "i1", "score": 0.9},
				{"item_id": "i2", "score": 0.8},
			},
			"user_profile": []string{"i3", "i4"},
		})
	})
	defer server.Close()

	result, err := client.RecommendProfile(1, []string{"i3", "i4"}, 2)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if len(result.Recommendations) != 2 {
		t.Fatalf("expected 2 recommendations, got %d", len(result.Recommendations))
	}
	if result.Recommendations[0].ItemId != "i1" {
		t.Errorf("expected first item_id i1, got %s", result.Recommendations[0].ItemId)
	}
	if len(result.UserProfile) != 2 {
		t.Fatalf("expected 2 user_profile items, got %d", len(result.UserProfile))
	}
	if result.UserProfile[0] != "i3" {
		t.Errorf("expected first user_profile item i3, got %s", result.UserProfile[0])
	}
}
