package niebulo

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (niebulo *Niebulo) onMessageReceived(message *tgbotapi.Message) {
	switch message.Command() {
	case "help", "h":
		niebulo.onHelpCommand(message)
	default:
		niebulo.sendTextOrPanick(message, "Unknown Command")
	}
}

func (niebulo *Niebulo) onHelpCommand(message *tgbotapi.Message) {
	niebulo.sendTextOrPanick(message, "Paste URL to audio you want to download from youtube")
}

func (niebulo *Niebulo) sendTextOrPanick(message *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	if _, err := niebulo.botApi.Send(msg); err != nil {
		panic(err)
	}
}
