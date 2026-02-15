package api

import (
	"context"
	"testing"
	"time"

	"recotem.org/cli/recotem/pkg/cfg"
)

func TestNewClient(t *testing.T) {
	config := cfg.RecotemConfig{Url: "http://localhost:8000"}
	client := NewClient(context.Background(), config)

	if client.Config.Url != "http://localhost:8000" {
		t.Errorf("expected URL http://localhost:8000, got %s", client.Config.Url)
	}
}

func TestNewApiClientWithApiKey(t *testing.T) {
	config := cfg.RecotemConfig{
		Url:    "http://localhost:8000",
		ApiKey: "test-api-key",
	}
	client := NewClient(context.Background(), config)

	apiClient, err := client.newApiClient()
	if err != nil {
		t.Fatalf("failed to create API client: %v", err)
	}
	if apiClient == nil {
		t.Fatal("expected non-nil API client")
	}
}

func TestNewApiClientWithJWT(t *testing.T) {
	config := cfg.RecotemConfig{
		Url:         "http://localhost:8000",
		AccessToken: "jwt-access-token",
	}
	client := NewClient(context.Background(), config)

	apiClient, err := client.newApiClient()
	if err != nil {
		t.Fatalf("failed to create API client: %v", err)
	}
	if apiClient == nil {
		t.Fatal("expected non-nil API client")
	}
}

func TestNewApiClientWithLegacyToken(t *testing.T) {
	config := cfg.RecotemConfig{
		Url:   "http://localhost:8000",
		Token: "legacy-token",
	}
	client := NewClient(context.Background(), config)

	apiClient, err := client.newApiClient()
	if err != nil {
		t.Fatalf("failed to create API client: %v", err)
	}
	if apiClient == nil {
		t.Fatal("expected non-nil API client")
	}
}

func TestNewUnauthenticatedClient(t *testing.T) {
	config := cfg.RecotemConfig{
		Url:   "http://localhost:8000",
		Token: "should-not-be-used",
	}
	client := NewClient(context.Background(), config)

	apiClient, err := client.newUnauthenticatedClient()
	if err != nil {
		t.Fatalf("failed to create unauthenticated client: %v", err)
	}
	if apiClient == nil {
		t.Fatal("expected non-nil unauthenticated client")
	}
}

func TestEnsureValidTokenWithApiKey(t *testing.T) {
	config := cfg.RecotemConfig{
		Url:    "http://localhost:8000",
		ApiKey: "test-key",
	}
	client := NewClient(context.Background(), config)

	// Should return nil immediately when API key is set
	err := client.EnsureValidToken()
	if err != nil {
		t.Errorf("expected nil error with API key, got %v", err)
	}
}

func TestEnsureValidTokenWithLegacyToken(t *testing.T) {
	config := cfg.RecotemConfig{
		Url:   "http://localhost:8000",
		Token: "legacy-token",
	}
	client := NewClient(context.Background(), config)

	// Should return nil immediately when legacy token is set
	err := client.EnsureValidToken()
	if err != nil {
		t.Errorf("expected nil error with legacy token, got %v", err)
	}
}

func TestEnsureValidTokenNoAuth(t *testing.T) {
	config := cfg.RecotemConfig{
		Url: "http://localhost:8000",
	}
	client := NewClient(context.Background(), config)

	err := client.EnsureValidToken()
	if err == nil {
		t.Error("expected error with no auth, got nil")
	}
}

func TestEnsureValidTokenNotExpired(t *testing.T) {
	future := time.Now().Add(10 * time.Minute)
	config := cfg.RecotemConfig{
		Url:         "http://localhost:8000",
		AccessToken: "valid-token",
		ExpiresAt:   &future,
	}
	client := NewClient(context.Background(), config)

	// Should return nil when token is not expired
	err := client.EnsureValidToken()
	if err != nil {
		t.Errorf("expected nil error with valid token, got %v", err)
	}
}

func TestRefreshTokenNoRefreshToken(t *testing.T) {
	config := cfg.RecotemConfig{
		Url:         "http://localhost:8000",
		AccessToken: "expired-token",
	}
	client := NewClient(context.Background(), config)

	_, err := client.RefreshToken()
	if err == nil {
		t.Error("expected error with no refresh token, got nil")
	}
}

func TestLoginResultStruct(t *testing.T) {
	now := time.Now()
	result := LoginResult{
		AccessToken:  "access",
		RefreshToken: "refresh",
		ExpiresAt:    now,
	}

	if result.AccessToken != "access" {
		t.Errorf("expected access token 'access', got %s", result.AccessToken)
	}
	if result.RefreshToken != "refresh" {
		t.Errorf("expected refresh token 'refresh', got %s", result.RefreshToken)
	}
	if !result.ExpiresAt.Equal(now) {
		t.Errorf("expected expires at %v, got %v", now, result.ExpiresAt)
	}
}
