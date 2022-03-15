package main

import (
	"fmt"
	"time"
)

/// dateTimeFormat sets the preferred format to YYYY/MM/DD
const dateTimeFormat = "2006-01-02"

type BotEvent struct {
	DateTime     time.Time
	RedditorName string
	Subreddit    string
	Event        string
	Permalink    string
}

type Logger interface {
	Log(event *BotEvent)
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

func (botLogger *BotLogger) Log(event *BotEvent) {
	var ymdNow = event.DateTime.Format(dateTimeFormat)
	var isoDateTime = event.DateTime.Format(time.RFC3339)

	err := botLogger.FileManager.write(
		botLogger.FilePath+ymdNow+".csv",
		[]string{
			isoDateTime,
			"u/" + event.RedditorName,
			"r/" + event.Subreddit,
			event.Event,
			event.Permalink,
		})
	if err != nil {
		// TODO: log to Sentry
		fmt.Errorf("failed to write file: %s", err)
	}
	fmt.Printf("%s, %s", ymdNow, event.Event)
}
