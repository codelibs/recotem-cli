package api

import (
	"net/http"
	"testing"
)

func TestDeleteTrainingDataSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			t.Errorf("expected DELETE method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusNoContent)
	})
	defer server.Close()

	err := client.DeleteTrainingData(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestGetTrainingDataSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, map[string]any{
			"count": 1,
			"results": []map[string]any{
				{"id": 1, "project": 1},
			},
		})
	})
	defer server.Close()

	result, err := client.GetTrainingData(nil, nil, nil, intPtr(1))
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
	if first.Project != 1 {
		t.Errorf("expected project 1, got %d", first.Project)
	}
}

func TestPreviewTrainingDataSuccess(t *testing.T) {
	previewBody := []byte("col1,col2\nval1,val2\n")
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET method, got %s", r.Method)
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(previewBody)
	})
	defer server.Close()

	data, err := client.PreviewTrainingData(1)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if string(data) != string(previewBody) {
		t.Errorf("expected preview body %q, got %q", string(previewBody), string(data))
	}
}
