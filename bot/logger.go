package main

import (
	"fmt"
	"time"
)

/// dateTimeFormat sets the preferred format to YYYY/MM/DD
const dateTimeFormat = "2006-01-02"
const eventDir = "events/"
const errDir = "errors/"

type BotEvent struct {
	DateTime     time.Time
	RedditorName string
	Subreddit    string
	Event        string
	Permalink    string
}

type Logger interface {
	Log(event *BotEvent)
	LogError(event *BotEvent)
}

type BotLogger struct {
	FilePath string
	FileManager
}

func NewBotLogger(filePath string) *BotLogger {
	return &BotLogger{
		filePath,
		csvFileManager{},
	}
}

func (l *BotLogger) Log(event *BotEvent) {
	writeLog(event, eventDir, l)
}

func (l *BotLogger) LogError(event *BotEvent) {
	writeLog(event, errDir, l)
}

func writeLog(event *BotEvent, dir string, l *BotLogger) {
	var ymdNow = event.DateTime.Format(dateTimeFormat)
	var isoDateTime = event.DateTime.Format(time.RFC3339)
	writeErr := l.FileManager.write(
		fmt.Sprintf("%s%s%s.csv", l.FilePath, dir, ymdNow),
		[]string{
			isoDateTime,
			fmt.Sprintf("u/%s", event.RedditorName),
			fmt.Sprintf("r/%s", event.Subreddit),
			event.Event,
			event.Permalink,
		})
	if writeErr != nil {
		// TODO: log to Sentry
		fmt.Printf("failed to write file: %s", writeErr)
	}
	fmt.Printf("%s, %s\n", ymdNow, event.Event)
}
