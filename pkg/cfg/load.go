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
		c := NewRectemConfig("http://localhost:8000")
		if saveErr := c.save(configPath); saveErr != nil {
			return RecotemConfig{}, saveErr
		}
		return c, nil
	}

	return NewRectemConfigFromFile(configPath)
}
