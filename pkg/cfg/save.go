package cfg

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func (c RecotemConfig) save(configPath string) error {
	out, err := yaml.Marshal(&c)
	if err != nil {
		return err
	}

	configDir := filepath.Dir(configPath)
	if _, statErr := os.Stat(configDir); os.IsNotExist(statErr) {
		if mkdirErr := os.MkdirAll(configDir, 0755); mkdirErr != nil {
			return mkdirErr
		}
	}

	err = os.WriteFile(configPath, out, 0600)
	if err != nil {
		return err
	}
	return nil
}

func SaveRecotemConfig(c RecotemConfig) error {
	configPath, err := configPath()
	if err != nil {
		return err
	}
	return c.save(configPath)
}
