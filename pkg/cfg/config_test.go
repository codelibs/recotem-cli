package cfg

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewRecotemConfig(t *testing.T) {
	url := "http://test.example.com"
	config := NewRecotemConfig(url)

	if config.Url != url {
		t.Errorf("expected URL %s, got %s", url, config.Url)
	}

	if config.Token != "" {
		t.Errorf("expected empty token, got %s", config.Token)
	}
}

func TestNewRecotemConfigFromFile(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create a test config file
	configPath := filepath.Join(tmpDir, "config.yaml")
	configContent := `url: http://test.example.com
token: test-token-123
`
	err = os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	// Test loading the config
	config, err := NewRecotemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if config.Url != "http://test.example.com" {
		t.Errorf("expected URL http://test.example.com, got %s", config.Url)
	}

	if config.Token != "test-token-123" {
		t.Errorf("expected token test-token-123, got %s", config.Token)
	}
}

func TestNewRecotemConfigFromFileNotFound(t *testing.T) {
	_, err := NewRecotemConfigFromFile("/nonexistent/path/config.yaml")
	if err == nil {
		t.Error("expected error for nonexistent file, got nil")
	}
}

func TestNewRecotemConfigFromFileInvalid(t *testing.T) {
	// Create a temporary directory
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create an invalid YAML file
	configPath := filepath.Join(tmpDir, "invalid.yaml")
	invalidContent := `this is not: valid: yaml: content`
	err = os.WriteFile(configPath, []byte(invalidContent), 0644)
	if err != nil {
		t.Fatalf("failed to write invalid config: %v", err)
	}

	_, err = NewRecotemConfigFromFile(configPath)
	if err == nil {
		t.Error("expected error for invalid YAML, got nil")
	}
}

func TestConfigPathConstant(t *testing.T) {
	if configFilename != ".recotem/config.yaml" {
		t.Errorf("expected configFilename to be .recotem/config.yaml, got %s", configFilename)
	}
}

func TestConfigPath(t *testing.T) {
	path, err := configPath()
	if err != nil {
		t.Fatalf("configPath() failed: %v", err)
	}

	// Should contain the config filename
	if !filepath.IsAbs(path) {
		t.Error("configPath should return an absolute path")
	}

	// Should end with .recotem/config.yaml
	if filepath.Base(path) != "config.yaml" {
		t.Errorf("path should end with config.yaml, got %s", path)
	}
}

// Test that RecotemConfig struct has the expected fields
func TestRecotemConfigStructure(t *testing.T) {
	config := RecotemConfig{
		Url:   "http://example.com",
		Token: "test-token",
	}

	if config.Url != "http://example.com" {
		t.Errorf("expected Url to be http://example.com, got %s", config.Url)
	}

	if config.Token != "test-token" {
		t.Errorf("expected Token to be test-token, got %s", config.Token)
	}
}

// Test with empty values
func TestNewRecotemConfigEmpty(t *testing.T) {
	config := NewRecotemConfig("")

	if config.Url != "" {
		t.Errorf("expected empty URL, got %s", config.Url)
	}

	if config.Token != "" {
		t.Errorf("expected empty token, got %s", config.Token)
	}
}

// Test loading config with only URL
func TestNewRecotemConfigFromFileOnlyUrl(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "config.yaml")
	configContent := `url: http://test.example.com
`
	err = os.WriteFile(configPath, []byte(configContent), 0644)
	if err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	config, err := NewRecotemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if config.Url != "http://test.example.com" {
		t.Errorf("expected URL http://test.example.com, got %s", config.Url)
	}

	if config.Token != "" {
		t.Errorf("expected empty token, got %s", config.Token)
	}
}

// Test loading empty config file
func TestNewRecotemConfigFromFileEmpty(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "recotem-test-*")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	configPath := filepath.Join(tmpDir, "empty.yaml")
	err = os.WriteFile(configPath, []byte(""), 0644)
	if err != nil {
		t.Fatalf("failed to write empty config: %v", err)
	}

	config, err := NewRecotemConfigFromFile(configPath)
	if err != nil {
		t.Fatalf("failed to load empty config: %v", err)
	}

	if config.Url != "" {
		t.Errorf("expected empty URL, got %s", config.Url)
	}

	if config.Token != "" {
		t.Errorf("expected empty token, got %s", config.Token)
	}
}
