package cfg

import (
	"io/ioutil"
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
	if _, err := os.Stat(configDir); os.IsNotExist(err) {
		if err := os.MkdirAll(configDir, 0755); err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(configPath, out, os.ModePerm)
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
