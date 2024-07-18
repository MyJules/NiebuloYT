package main

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type BotConfig struct {
	Telegram_token string
}

func BotConfigFromYamlFile(configPath string) (botConfig BotConfig, err error) {
	resultConfig := BotConfig{}

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
