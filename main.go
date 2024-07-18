package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	botConfig, err := BotConfigFromYamlFile("config.yaml")
	if err != nil {
		panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(botConfig.Telegram_token)
	if err != nil {
		panic(err)
	}

	bot.Debug = false
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		image := tgbotapi.FileURL("https://img.freepik.com/free-photo/cute-ai-generated-cartoon-bunny_23-2150288869.jpg?t=st=1721247939~exp=1721251539~hmac=c1e324a46d93b152d723b22c8c5b2c5f0339783c315818fa69f83a90136d5ec4&w=740")
		photo := tgbotapi.NewPhoto(update.Message.Chat.ID, image)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(photo); err != nil {
			panic(err)
		}

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
