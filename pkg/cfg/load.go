package cfg

import (
	"os"
)

func LoadRecotemConfig() (RecotemConfig, error) {
	configPath, err := configPath()
	if err != nil {
		return RecotemConfig{}, err
	}

	_, err = os.Stat(configPath)
	if err != nil {
		c := NewRecotemConfig("http://localhost:8000")
		if saveErr := c.save(configPath); saveErr != nil {
			return RecotemConfig{}, saveErr
		}
		return c, nil
	}

	config, err := NewRecotemConfigFromFile(configPath)
	if err != nil {
		return RecotemConfig{}, err
	}

	// Migrate legacy config: if Token is set but no JWT tokens, keep it for backward compat
	return config, nil
}
