package api

import (
	"net/http"
	"testing"
)

func TestGetApiKeysSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"count": 1,
			"results": []map[string]any{
				{"id": 1, "name": "key1", "key": "ak-xxx", "is_active": true},
			},
		})
	})
	defer server.Close()

	result, err := client.GetApiKeys(nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Count == nil || *result.Count != 1 {
		t.Fatalf("expected count 1, got %v", result.Count)
	}
	if result.Results == nil || len(*result.Results) != 1 {
		t.Fatal("expected 1 result")
	}
	key := (*result.Results)[0]
	if key.Name != "key1" {
		t.Errorf("expected name 'key1', got %s", key.Name)
	}
	if key.IsActive != true {
		t.Errorf("expected is_active true, got %v", key.IsActive)
	}
}

func TestCreateApiKeySuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusCreated, map[string]any{
			"id":        1,
			"name":      "new-key",
			"key":       "ak-newkey123",
			"is_active": true,
		})
	})
	defer server.Close()

	result, err := client.CreateApiKey("new-key")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Name != "new-key" {
		t.Errorf("expected name 'new-key', got %s", result.Name)
	}
	if result.Key == nil || *result.Key != "ak-newkey123" {
		t.Errorf("expected key 'ak-newkey123', got %v", result.Key)
	}
	if result.IsActive != true {
		t.Errorf("expected is_active true, got %v", result.IsActive)
	}
}

func TestGetApiKeySuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"name":      "key1",
			"key":       "ak-xxx",
			"is_active": true,
		})
	})
	defer server.Close()

	result, err := client.GetApiKey(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.Id == nil || *result.Id != 1 {
		t.Errorf("expected id 1, got %v", result.Id)
	}
	if result.Name != "key1" {
		t.Errorf("expected name 'key1', got %s", result.Name)
	}
}

func TestDeleteApiKeySuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteApiKey(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestRevokeApiKeySuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusOK, map[string]any{
			"id":        1,
			"name":      "key1",
			"is_active": false,
		})
	})
	defer server.Close()

	err := client.RevokeApiKey(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
