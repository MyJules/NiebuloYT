package niebulo

import (
	"net/url"

	yt "github.com/MyJules/NiebuloYT/yt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (niebulo *Niebulo) onMessageReceived(message *tgbotapi.Message) {
	niebulo.logger.Info("Message received: " + message.Text)

	//Handle url
	url, err := url.ParseRequestURI(message.Text)
	if err != nil {
		niebulo.sendTextOrPanic(message, "Paste send URL to audio you want to download from youtube")
		return
	}

	ytAudio := yt.NewYTAudio(url)
	ytAudio.DownloadAudio()
	ytAudio.RegisterOnAudioReady(func() {
		file := tgbotapi.FilePath(ytAudio.GetAudioFilePath())

		niebulo.logger.Info(ytAudio.GetAudioFilePath())

		msg := tgbotapi.NewAudio(message.Chat.ID, file)
		if _, err := niebulo.botApi.Send(msg); err != nil {
			niebulo.logger.Error("Failed to send message")
			panic(err)
		}
		ytAudio.ClearAudio()
	})
}

func (niebulo *Niebulo) sendTextOrPanic(message *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	if _, err := niebulo.botApi.Send(msg); err != nil {
		niebulo.logger.Error("Failed to send message")
		panic(err)
	}
	niebulo.logger.Info("Message sent: " + text)
}
