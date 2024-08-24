package niebulo

import (
	"net/url"

	yt "github.com/MyJules/NiebuloYT/yt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (niebulo *Niebulo) onMessageReceived(message *tgbotapi.Message) {
	niebulo.logger.Info("Message received: " + message.Text)
	niebulo.logger.Info("Message rceived from: " + message.From.UserName)
	messageUserName := message.From.UserName

	//Handle url
	url, err := url.ParseRequestURI(message.Text)
	if err != nil {
		niebulo.sendTextOrPanic(message, "Paste send URL to audio you want to download from youtube")
		return
	}

	_, exists := niebulo.taskMap[messageUserName]
	if exists {
		niebulo.sendTextOrPanic(message, "Please wait for audio to download")
		return
	}

	//Handle downloading youtube audio
	ytAudio := yt.NewYTAudio(url)
	ytAudio.RegisterOnAudioFailed(func() {
		delete(niebulo.taskMap, messageUserName)
		niebulo.sendTextOrPanic(message, "Failed to download audio from youtube")
	})
	ytAudio.RegisterOnAudioReady(func() {
		delete(niebulo.taskMap, messageUserName)
		niebulo.sendTextOrPanic(message, "Here is your audio:")
		niebulo.sendAudioOrPanic(message, &ytAudio)
		ytAudio.ClearAudio()
	})
	niebulo.taskMap[messageUserName] = &ytAudio
	ytAudio.DownloadAudio()
	niebulo.sendTextOrPanic(message, "Started downloading audio")
}

func (niebulo *Niebulo) sendTextOrPanic(message *tgbotapi.Message, text string) {
	msg := tgbotapi.NewMessage(message.Chat.ID, text)

	if _, err := niebulo.botApi.Send(msg); err != nil {
		niebulo.logger.Error("Failed to send message")
		panic(err)
	}
	niebulo.logger.Info("Message sent: " + text)
}

func (niebulo *Niebulo) sendAudioOrPanic(message *tgbotapi.Message, ytAudio *yt.YTAudio) {
	file := tgbotapi.FilePath(ytAudio.GetAudioFilePath())
	msg := tgbotapi.NewAudio(message.Chat.ID, file)
	if _, err := niebulo.botApi.Send(msg); err != nil {
		niebulo.logger.Error("Failed to send message")
		panic(err)
	}
}
