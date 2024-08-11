package yt

import (
	"bytes"
	"net/url"
	"os"
	"os/exec"
	"regexp"
	"sync"

	"go.uber.org/zap"
)

type NiebuloYTState int32

const (
	Loading NiebuloYTState = iota
	Idle
	Ready
	Failed
)

type YTAudio struct {
	state         NiebuloYTState
	stateMutex    sync.Mutex
	url           *url.URL
	logger        *zap.Logger
	audioPath     string
	onAudioReady  func()
	onAudioFailed func()
}

func NewYTAudio(youtubeURL *url.URL) YTAudio {
	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	logger.Info("YTAudio created with url: " + youtubeURL.String())
	return YTAudio{state: Idle, stateMutex: sync.Mutex{}, url: youtubeURL, logger: logger, audioPath: "", onAudioReady: nil, onAudioFailed: nil}
}

func (ytAudio *YTAudio) RegisterOnAudioReady(fn func()) {
	ytAudio.onAudioReady = fn
}

func (ytAudio *YTAudio) RegisterOnAudioFailed(fn func()) {
	ytAudio.onAudioFailed = fn
}

func (ytAudio *YTAudio) GetAudioFilePath() string {
	if ytAudio.state != Ready {
		ytAudio.logger.Warn("Trying to get file audio path on unready type")
		return ytAudio.audioPath
	}

	return ytAudio.audioPath
}

func (ytAudio *YTAudio) DownloadAudio() {
	if ytAudio.state == Loading || ytAudio.state == Ready {
		return
	}

	go func() {
		ytAudio.logger.Info("Started downloading audio")
		ytAudio.setState(Loading)

		//download audio here
		stdout := new(bytes.Buffer)
		stderr := new(bytes.Buffer)
		ytdl := exec.Command("../bin/yt-dlp.exe", "-x", "--audio-format", "mp3", "-o", "%(title)s", ytAudio.url.String())
		ytdl.Stdout = stdout
		ytdl.Stderr = stderr
		err := ytdl.Run()
		if err != nil {
			ytAudio.logger.Error("Failed to download youtube audio")
			ytAudio.logger.Error(err.Error())
			ytAudio.logger.Error(stderr.String())
			ytAudio.setState(Failed)
			if ytAudio.onAudioFailed != nil {
				ytAudio.onAudioFailed()
			}
			return
		}

		audioFileName := extractDestination(stdout.String()) + ".mp3"
		ytAudio.logger.Info("Saved audio to file: " + audioFileName)

		ytAudio.audioPath = audioFileName
		ytAudio.logger.Info("Finished downloading audio")
		ytAudio.setState(Ready)
		if ytAudio.onAudioReady != nil {
			ytAudio.onAudioReady()
		}
	}()
}

func (ytAudio *YTAudio) ClearAudio() {
	e := os.Remove(ytAudio.audioPath)
	if e != nil {
		ytAudio.logger.Warn("Failed to clear audio")
		return
	}
	ytAudio.setState(Idle)
}

func (ytAudio *YTAudio) setState(newState NiebuloYTState) {
	ytAudio.stateMutex.Lock()
	ytAudio.state = newState
	ytAudio.stateMutex.Unlock()
}

func extractDestination(text string) string {
	re := regexp.MustCompile(`Destination:\s*(.*)`)
	match := re.FindStringSubmatch(text)
	if len(match) == 0 {
		return ""
	}
	return match[1]
}
