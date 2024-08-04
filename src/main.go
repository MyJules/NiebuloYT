package main

import (
	"github.com/MyJules/NiebuloYT/niebulo"
)

func main() {
	niebuloConfig, err := niebulo.BotConfigFromYamlFile("../config/config.yaml")
	if err != nil {
		panic(err)
	}

	niebulo := niebulo.NewNiebuloBot(niebuloConfig)
	defer niebulo.Delete()
	niebulo.Start()
}
