package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"testing"

	"recotem.org/cli/recotem/pkg/cfg"
)

func TestLoginSuccess(t *testing.T) {
	server, client := newTestServerUnauthenticated(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/api/v1/auth/login/") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		// Verify request body contains username and password
		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode request body: %v", err)
		}
		if body["username"] != "testuser" {
			t.Errorf("expected username 'testuser', got %v", body["username"])
		}
		if body["password"] != "testpass" {
			t.Errorf("expected password 'testpass', got %v", body["password"])
		}

		jsonResponse(w, http.StatusOK, map[string]interface{}{
			"access_token":  "abc",
			"refresh_token": "xyz",
			"user": map[string]interface{}{
				"username": "testuser",
			},
		})
	})
	defer server.Close()

	result, err := client.Login("testuser", "testpass")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.AccessToken != "abc" {
		t.Errorf("expected access token 'abc', got %s", result.AccessToken)
	}
	if result.RefreshToken != "xyz" {
		t.Errorf("expected refresh token 'xyz', got %s", result.RefreshToken)
	}
	if result.ExpiresAt.IsZero() {
		t.Error("expected non-zero ExpiresAt")
	}
}

func TestLoginInvalidCredentials(t *testing.T) {
	server, client := newTestServerUnauthenticated(func(w http.ResponseWriter, r *http.Request) {
		jsonResponse(w, http.StatusUnauthorized, map[string]string{
			"detail": "No active account found with the given credentials",
		})
	})
	defer server.Close()

	result, err := client.Login("baduser", "badpass")
	if err == nil {
		t.Fatal("expected error for invalid credentials, got nil")
	}
	if result != nil {
		t.Errorf("expected nil result, got %v", result)
	}
	if !strings.Contains(err.Error(), "401") {
		t.Errorf("expected error to contain '401', got %s", err.Error())
	}
}

func TestLogoutSuccess(t *testing.T) {
	server, client := newTestServer(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/api/v1/auth/token/blacklist/") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		w.WriteHeader(http.StatusOK)
	})
	defer server.Close()

	// Set a refresh token so Logout doesn't return early
	client.Config.RefreshToken = "test-refresh-token"

	err := client.Logout()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}

func TestLogoutNoRefreshToken(t *testing.T) {
	// Create a client with no refresh token
	client := NewClient(context.Background(), cfg.RecotemConfig{
		Url:         "http://localhost:9999",
		AccessToken: "test-token",
	})

	// Should return nil immediately when no refresh token is set
	err := client.Logout()
	if err != nil {
		t.Fatalf("expected nil error for no refresh token, got %v", err)
	}
}

func TestRefreshTokenSuccess(t *testing.T) {
	server, client := newTestServerUnauthenticated(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("expected POST, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/api/v1/auth/token/refresh/") {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		jsonResponse(w, http.StatusOK, map[string]string{
			"access": "new-access",
		})
	})
	defer server.Close()

	// RefreshToken requires a refresh token to be set
	client.Config.RefreshToken = "existing-refresh-token"

	result, err := client.RefreshToken()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if result == nil {
		t.Fatal("expected non-nil result")
	}
	if result.AccessToken != "new-access" {
		t.Errorf("expected access token 'new-access', got %s", result.AccessToken)
	}
	if result.RefreshToken != "existing-refresh-token" {
		t.Errorf("expected refresh token 'existing-refresh-token', got %s", result.RefreshToken)
	}
	if result.ExpiresAt.IsZero() {
		t.Error("expected non-zero ExpiresAt")
	}
}
