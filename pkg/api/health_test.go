package api

import (
	"net/http"
	"strings"
	"testing"
)

func TestPingSuccess(t *testing.T) {
	server, client := newTestServerUnauthenticated(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/api/v1/ping/") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		jsonResponse(w, http.StatusOK, map[string]string{
			"status": "ok",
		})
	})
	defer server.Close()

	body, err := client.Ping()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if body == nil {
		t.Fatal("expected non-nil body")
	}
	if !strings.Contains(string(body), `"status":"ok"`) {
		t.Errorf("expected body to contain '{\"status\":\"ok\"}', got %s", string(body))
	}
}

func TestPingServerError(t *testing.T) {
	server, client := newTestServerUnauthenticated(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusInternalServerError, map[string]string{
			"detail": "Internal Server Error",
		})
	})
	defer server.Close()

	body, err := client.Ping()
	if err == nil {
		t.Fatal("expected error for server error, got nil")
	}
	if body != nil {
		t.Errorf("expected nil body, got %v", body)
	}
	if !strings.Contains(err.Error(), "500") {
		t.Errorf("expected error to contain '500', got %s", err.Error())
	}
}
