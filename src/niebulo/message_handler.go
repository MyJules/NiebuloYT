package niebulo

import (
	"net/url"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (niebulo *Niebulo) onMessageReceived(message *tgbotapi.Message) {
	niebulo.logger.Info("Message received: " + message.Text)

	//Handle commands
	switch message.Command() {
	case "help", "h":
		niebulo.onHelpCommand(message)
	}

	//Handle url
	_, err := url.ParseRequestURI(message.Text)
	if err != nil {
		niebulo.sendTextOrPanic(message, "Paste send URL to audio you want to download from youtube")
	}
}

func (niebulo *Niebulo) onHelpCommand(message *tgbotapi.Message) {
	niebulo.sendTextOrPanic(message, "Paste send URL to audio you want to download from youtube")
}

func (niebulo *Niebulo) sendTextOrPanic(message *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	if _, err := niebulo.botApi.Send(msg); err != nil {
		niebulo.logger.Error("Failed to send message")
		panic(err)
	}
	niebulo.logger.Info("Message sent: " + text)
}
