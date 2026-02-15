package api

import (
	"net/http"
	"testing"

	"recotem.org/cli/recotem/pkg/openapi"
)

func TestCreateSplitConfigSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":              1,
			"name":            "split1",
			"scheme":          "RG",
			"heldout_ratio":   0.2,
			"n_heldout":       nil,
			"test_user_ratio": 0.5,
			"n_test_users":    nil,
			"random_seed":     42,
		})
	})
	defer server.Close()

	scheme := openapi.SchemeEnum("RG")
	heldoutRatio := float32(0.2)
	testUserRatio := float32(0.5)
	result, err := client.CreateSplitConfig(stringPtr("split1"), &scheme, &heldoutRatio, nil, &testUserRatio, nil, intPtr(42))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name == nil || *result.Name != "split1" {
		t.Errorf("expected name 'split1', got %v", result.Name)
	}
	if result.Scheme == nil || *result.Scheme != openapi.SchemeEnum("RG") {
		t.Errorf("expected scheme 'RG', got %v", result.Scheme)
	}
	if result.RandomSeed == nil || *result.RandomSeed != 42 {
		t.Errorf("expected random_seed 42, got %v", result.RandomSeed)
	}
}

func TestDeleteSplitConfigSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteSplitConfig(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetSplitConfigsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, []map[string]any{
			{
				"id":     1,
				"name":   "split1",
				"scheme": "RG",
			},
			{
				"id":     2,
				"name":   "split2",
				"scheme": "TG",
			},
		})
	})
	defer server.Close()

	configs, err := client.GetSplitConfigs(nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if configs == nil {
		t.Fatal("expected non-nil configs")
	}
	if len(*configs) != 2 {
		t.Fatalf("expected 2 configs, got %d", len(*configs))
	}
	if (*configs)[0].Id == nil || *(*configs)[0].Id != 1 {
		t.Errorf("expected first id 1, got %v", (*configs)[0].Id)
	}
	if (*configs)[0].Name == nil || *(*configs)[0].Name != "split1" {
		t.Errorf("expected first name 'split1', got %v", (*configs)[0].Name)
	}
	if (*configs)[1].Id == nil || *(*configs)[1].Id != 2 {
		t.Errorf("expected second id 2, got %v", (*configs)[1].Id)
	}
}

func TestUpdateSplitConfigSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":          1,
			"name":        "split1-updated",
			"scheme":      "TG",
			"random_seed": 99,
		})
	})
	defer server.Close()

	scheme := openapi.SchemeEnum("TG")
	result, err := client.UpdateSplitConfig(1, stringPtr("split1-updated"), &scheme, nil, nil, nil, nil, intPtr(99))
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name == nil || *result.Name != "split1-updated" {
		t.Errorf("expected name 'split1-updated', got %v", result.Name)
	}
	if result.Scheme == nil || *result.Scheme != openapi.SchemeEnum("TG") {
		t.Errorf("expected scheme 'TG', got %v", result.Scheme)
	}
	if result.RandomSeed == nil || *result.RandomSeed != 99 {
		t.Errorf("expected random_seed 99, got %v", result.RandomSeed)
	}
}
