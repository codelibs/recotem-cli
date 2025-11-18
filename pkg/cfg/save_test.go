package cfg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveRecotemConfig(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test config
	config := RecotemConfig{
		Url:   "http://test.example.com",
		Token: "test-token-123",
	}

	// Test save to specific path
	configPath := filepath.Join(tmpDir, ".recotem", "config.yaml")
	err = config.save(configPath)
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Verify file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("config file was not created at %s", configPath)
	}

	// Load the saved config and verify
	loadedConfig, err := NewRectemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load saved config: %v", err)
	}

	if loadedConfig.Url != config.Url {
		t.Errorf("expected URL %s, got %s", config.Url, loadedConfig.Url)
	}

	if loadedConfig.Token != config.Token {
		t.Errorf("expected token %s, got %s", config.Token, loadedConfig.Token)
	}
}

func TestSaveCreatesDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create test config
	config := RecotemConfig{
		Url:   "http://test.example.com",
		Token: "test-token-456",
	}

	// Test save to a path where directory doesn't exist
	configPath := filepath.Join(tmpDir, "new", "directory", "config.yaml")
	err = config.save(configPath)
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Verify directory was created
	configDir := filepath.Dir(configPath)
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		t.Errorf("config directory was not created at %s", configDir)
	}

	// Verify file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		t.Errorf("config file was not created at %s", configPath)
	}
}

func TestSaveOverwritesExisting(t *testing.T) {
	// Create a temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "config.yaml")

	// Save first config
	config1 := RecotemConfig{
		Url:   "http://first.example.com",
		Token: "first-token",
	}
	err = config1.save(configPath)
	if err != nil {
		t.Fatalf("failed to save first config: %v", err)
	}

	// Save second config to same path
	config2 := RecotemConfig{
		Url:   "http://second.example.com",
		Token: "second-token",
	}
	err = config2.save(configPath)
	if err != nil {
		t.Fatalf("failed to save second config: %v", err)
	}

	// Load and verify it has the second config
	loadedConfig, err := NewRectemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loadedConfig.Url != config2.Url {
		t.Errorf("expected URL %s, got %s", config2.Url, loadedConfig.Url)
	}

	if loadedConfig.Token != config2.Token {
		t.Errorf("expected token %s, got %s", config2.Token, loadedConfig.Token)
	}
}

func TestSaveEmptyConfig(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create empty config
	config := RecotemConfig{}
	configPath := filepath.Join(tmpDir, "empty.yaml")

	err = config.save(configPath)
	if err != nil {
		t.Fatalf("failed to save empty config: %v", err)
	}

	// Load and verify
	loadedConfig, err := NewRectemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load empty config: %v", err)
	}

	if loadedConfig.Url != "" {
		t.Errorf("expected empty URL, got %s", loadedConfig.Url)
	}

	if loadedConfig.Token != "" {
		t.Errorf("expected empty token, got %s", loadedConfig.Token)
	}
}

func TestSaveWithSpecialCharacters(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create config with special characters
	config := RecotemConfig{
		Url:   "http://test.example.com:8080/api/v1",
		Token: "token-with-special-chars-!@#$%",
	}
	configPath := filepath.Join(tmpDir, "special.yaml")

	err = config.save(configPath)
	if err != nil {
		t.Fatalf("failed to save config with special chars: %v", err)
	}

	// Load and verify
	loadedConfig, err := NewRectemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if loadedConfig.Url != config.Url {
		t.Errorf("expected URL %s, got %s", config.Url, loadedConfig.Url)
	}

	if loadedConfig.Token != config.Token {
		t.Errorf("expected token %s, got %s", config.Token, loadedConfig.Token)
	}
}

// Test file permissions
func TestSaveFilePermissions(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	config := RecotemConfig{
		Url:   "http://test.example.com",
		Token: "test-token",
	}
	configPath := filepath.Join(tmpDir, "config.yaml")

	err = config.save(configPath)
	if err != nil {
		t.Fatalf("failed to save config: %v", err)
	}

	// Check file permissions
	fileInfo, err := os.Stat(configPath)
	if err != nil {
		t.Fatalf("failed to stat config file: %v", err)
	}

	// os.ModePerm is 0777, so the file should have those permissions (modified by umask)
	mode := fileInfo.Mode()
	if mode&0400 == 0 {
		t.Error("config file should be readable by owner")
	}
	if mode&0200 == 0 {
		t.Error("config file should be writable by owner")
	}
}
