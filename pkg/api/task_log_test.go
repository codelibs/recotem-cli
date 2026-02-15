package api

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetTaskLogsSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/task-log") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		jsonResponse(w, http.StatusOK, []map[string]any{
			{
				"id":       1,
				"task":     1,
				"contents": "Training completed",
			},
		})
	})
	defer server.Close()

	result, err := client.GetTaskLogs(nil, nil, nil)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if len(result) == 0 {
		t.Fatal("expected non-empty response body")
	}
	body := string(result)
	if !strings.Contains(body, "Training completed") {
		t.Errorf("expected body to contain 'Training completed', got %s", body)
	}
}

func TestGetTaskLogsError(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasPrefix(r.URL.Path, "/api/v1/task-log") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"detail":"internal server error"}`))
	})
	defer server.Close()

	result, err := client.GetTaskLogs(nil, nil, nil)
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if result != nil {
		t.Errorf("expected nil result on error, got %v", result)
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("expected error to contain '500', got %s", err.Error())
	}
}
