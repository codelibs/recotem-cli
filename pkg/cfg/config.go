package cfg

import (
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

const (
	configFilename = ".recotem/config.yaml"
)

type RecotemConfig struct {
	Url          string     `yaml:"url"`
	Token        string     `yaml:"token,omitempty"`
	AccessToken  string     `yaml:"access_token,omitempty"`
	RefreshToken string     `yaml:"refresh_token,omitempty"`
	ExpiresAt    *time.Time `yaml:"expires_at,omitempty"`
	ApiKey       string     `yaml:"api_key,omitempty"`
}

func NewRecotemConfig(url string) RecotemConfig {
	return RecotemConfig{Url: url}
}

func NewRecotemConfigFromFile(filename string) (RecotemConfig, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return RecotemConfig{}, err
	}

	config := RecotemConfig{}
	err = yaml.Unmarshal(buf, &config)
	if err != nil {
		return RecotemConfig{}, err
	}

	return config, nil
}

func configPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return home + string(os.PathSeparator) + configFilename, nil
}

// ClearTokens removes all authentication tokens from the config
func (c *RecotemConfig) ClearTokens() {
	c.Token = ""
	c.AccessToken = ""
	c.RefreshToken = ""
	c.ExpiresAt = nil
}

// IsTokenExpired checks if the access token is expired or will expire within 30 seconds
func (c *RecotemConfig) IsTokenExpired() bool {
	if c.ExpiresAt == nil {
		return true
	}
	return time.Now().Add(30 * time.Second).After(*c.ExpiresAt)
}
