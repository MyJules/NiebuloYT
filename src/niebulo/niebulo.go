package niebulo

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Niebulo struct {
	botApi *tgbotapi.BotAPI
	logger *zap.Logger
}

func NewNiebuloBot(niebuloConfig NiebuloConfig) Niebulo {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(niebuloConfig.Telegram_token)
	if err != nil {
		panic(err)
	}

	bot.Debug = false

	logger.Info("Niebulo Created")
	return Niebulo{botApi: bot, logger: logger}
}

func (niebulo *Niebulo) Delete() {
	niebulo.logger.Info("Niebulo Deleted")
	niebulo.logger.Sync()
}

func (niebulo *Niebulo) Start() {
	niebulo.logger.Info("Niebulo Started")

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := niebulo.botApi.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go niebulo.onMessageReceived(update.Message)
	}
}
