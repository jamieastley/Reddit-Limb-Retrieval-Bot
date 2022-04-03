package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
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

type BotLogger interface {
	Console
	LogEvent(event *BotEvent)
	LogError(event *BotEvent)
}

type Console interface {
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
}

type botLogger struct {
	FilePath string
	FileManager
	*consoleLogger
}

type consoleLogger struct {
	logger *zap.SugaredLogger
}

func NewBotLogger(filePath string) *botLogger {
	encoder := getEncoder()
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, getLogWriter(), zapcore.DebugLevel),
		zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel),
	)
	z := zap.New(core, zap.AddCaller())
	logger := z.Sugar()
	defer logger.Sync()

	return &botLogger{
		filePath,
		csvFileManager{},
		&consoleLogger{logger},
	}
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/debug/debug.log", loggingDirPath),
		MaxSize:    1,
		MaxBackups: 30,
		MaxAge:     60,
		Compress:   true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (l *botLogger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l *botLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *botLogger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

func (l *botLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *botLogger) LogEvent(event *BotEvent) {
	writeLog(event, eventDir, l)
}

func (l *botLogger) LogError(event *BotEvent) {
	writeLog(event, errDir, l)
}

func writeLog(event *BotEvent, dir string, l *botLogger) {
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
		},
	)
	if writeErr != nil {
		// TODO: log to Sentry
		l.logger.Error(writeErr.Error())
	} else {
		l.logger.Debug("Successfully wrote to bot logs: ", event)
	}

}
