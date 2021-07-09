package cfg

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFilename = ".recotem/config.yaml"
)

type RecotemConfig struct {
	Url   string
	Token string
}

func NewRectemConfig(url string) (c RecotemConfig) {
	c = RecotemConfig{}
	c.Url = url
	return
}

func NewRectemConfigFromFile(filename string) (RecotemConfig, error) {
	buf, err := ioutil.ReadFile(filename)
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
