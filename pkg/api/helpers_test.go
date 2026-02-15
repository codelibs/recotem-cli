package api

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"recotem.org/cli/recotem/pkg/cfg"
)

// newTestServer creates a httptest server and a Client pointing to it.
func newTestServer(handler http.HandlerFunc) (*httptest.Server, Client) {
	server := httptest.NewServer(handler)
	config := cfg.RecotemConfig{
		Url:         server.URL,
		AccessToken: "test-token",
	}
	client := NewClient(context.Background(), config)
	return server, client
}

// newTestServerUnauthenticated creates a httptest server and a Client with no auth tokens.
func newTestServerUnauthenticated(handler http.HandlerFunc) (*httptest.Server, Client) {
	server := httptest.NewServer(handler)
	config := cfg.RecotemConfig{
		Url: server.URL,
	}
	client := NewClient(context.Background(), config)
	return server, client
}

// jsonResponse writes a JSON response with the given status code and body.
func jsonResponse(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

// intPtr returns a pointer to an int.
func intPtr(i int) *int {
	return &i
}

// stringPtr returns a pointer to a string.
func stringPtr(s string) *string {
	return &s
}
