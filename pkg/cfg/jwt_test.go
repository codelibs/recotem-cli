package cfg

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestClearTokens(t *testing.T) {
	future := time.Now().Add(10 * time.Minute)
	config := RecotemConfig{
		Url:          "http://example.com",
		Token:        "legacy-token",
		AccessToken:  "access-token",
		RefreshToken: "refresh-token",
		ExpiresAt:    &future,
	}

	config.ClearTokens()

	if config.Token != "" {
		t.Errorf("expected empty Token, got %s", config.Token)
	}
	if config.AccessToken != "" {
		t.Errorf("expected empty AccessToken, got %s", config.AccessToken)
	}
	if config.RefreshToken != "" {
		t.Errorf("expected empty RefreshToken, got %s", config.RefreshToken)
	}
	if config.ExpiresAt != nil {
		t.Errorf("expected nil ExpiresAt, got %v", config.ExpiresAt)
	}
	// Url should be preserved
	if config.Url != "http://example.com" {
		t.Errorf("expected Url to be preserved, got %s", config.Url)
	}
}

func TestIsTokenExpiredNilExpiry(t *testing.T) {
	config := RecotemConfig{
		AccessToken: "some-token",
		ExpiresAt:   nil,
	}

	if !config.IsTokenExpired() {
		t.Error("expected expired with nil ExpiresAt")
	}
}

func TestIsTokenExpiredFutureExpiry(t *testing.T) {
	future := time.Now().Add(10 * time.Minute)
	config := RecotemConfig{
		AccessToken: "some-token",
		ExpiresAt:   &future,
	}

	if config.IsTokenExpired() {
		t.Error("expected not expired with future ExpiresAt")
	}
}

func TestIsTokenExpiredPastExpiry(t *testing.T) {
	past := time.Now().Add(-10 * time.Minute)
	config := RecotemConfig{
		AccessToken: "some-token",
		ExpiresAt:   &past,
	}

	if !config.IsTokenExpired() {
		t.Error("expected expired with past ExpiresAt")
	}
}

func TestIsTokenExpiredWithinGracePeriod(t *testing.T) {
	// Token expires in 20 seconds, but grace period is 30 seconds
	nearFuture := time.Now().Add(20 * time.Second)
	config := RecotemConfig{
		AccessToken: "some-token",
		ExpiresAt:   &nearFuture,
	}

	if !config.IsTokenExpired() {
		t.Error("expected expired within 30-second grace period")
	}
}

func TestJWTConfigRoundTrip(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	expiresAt := time.Now().Add(5 * time.Minute).Truncate(time.Second).UTC()
	original := RecotemConfig{
		Url:          "http://example.com",
		AccessToken:  "access-abc123",
		RefreshToken: "refresh-xyz789",
		ExpiresAt:    &expiresAt,
		ApiKey:       "ak-test-key",
	}

	configPath := filepath.Join(tmpDir, "jwt-config.yaml")
	err = original.save(configPath)
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	loaded, err := NewRecotemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loaded.AccessToken != original.AccessToken {
		t.Errorf("AccessToken mismatch: expected %s, got %s", original.AccessToken, loaded.AccessToken)
	}
	if loaded.RefreshToken != original.RefreshToken {
		t.Errorf("RefreshToken mismatch: expected %s, got %s", original.RefreshToken, loaded.RefreshToken)
	}
	if loaded.ApiKey != original.ApiKey {
		t.Errorf("ApiKey mismatch: expected %s, got %s", original.ApiKey, loaded.ApiKey)
	}
	if loaded.ExpiresAt == nil {
		t.Fatal("expected non-nil ExpiresAt")
	}
	if !loaded.ExpiresAt.Equal(expiresAt) {
		t.Errorf("ExpiresAt mismatch: expected %v, got %v", expiresAt, *loaded.ExpiresAt)
	}
}

func TestLegacyConfigCompatibility(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Simulate a legacy config file with only url and token
	configPath := filepath.Join(tmpDir, "legacy.yaml")
	legacyContent := `url: http://legacy.example.com
token: legacy-token-123
`
	err = os.WriteFile(configPath, []byte(legacyContent), 0644)
	if err != nil {
		t.Fatalf("failed to write legacy config: %v", err)
	}

	config, err := NewRecotemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load legacy config: %v", err)
	}

	if config.Url != "http://legacy.example.com" {
		t.Errorf("expected legacy URL, got %s", config.Url)
	}
	if config.Token != "legacy-token-123" {
		t.Errorf("expected legacy token, got %s", config.Token)
	}
	// JWT fields should be empty/nil
	if config.AccessToken != "" {
		t.Errorf("expected empty AccessToken, got %s", config.AccessToken)
	}
	if config.RefreshToken != "" {
		t.Errorf("expected empty RefreshToken, got %s", config.RefreshToken)
	}
	if config.ExpiresAt != nil {
		t.Errorf("expected nil ExpiresAt, got %v", config.ExpiresAt)
	}
}
