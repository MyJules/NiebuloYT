package niebulo

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type Niebulo struct {
	botApi *tgbotapi.BotAPI
}

func NewNiebuloBot(niebuloConfig NiebuloConfig) Niebulo {
	bot, err := tgbotapi.NewBotAPI(niebuloConfig.Telegram_token)
	if err != nil {
		panic(err)
	}

	return Niebulo{botApi: bot}
}

func (niebulo Niebulo) Start() {
	niebulo.botApi.Debug = false
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := niebulo.botApi.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := niebulo.botApi.Send(msg); err != nil {
			panic(err)
		}
	}
}
