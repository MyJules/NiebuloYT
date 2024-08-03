package niebulo

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type NiebuloConfig struct {
	Telegram_token string
}

func BotConfigFromYamlFile(configPath string) (botConfig NiebuloConfig, err error) {
	resultConfig := NiebuloConfig{}

	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		return resultConfig, err
	}

	err = yaml.Unmarshal(yamlFile, &resultConfig)
	if err != nil {
		return resultConfig, err
	}

	return resultConfig, nil
}
