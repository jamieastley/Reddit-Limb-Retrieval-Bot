package main

import (
	"github.com/turnage/graw/reddit"
	"regexp"
	"time"
)

type TypePrefix string

const (
	Comment   TypePrefix = "t1_"
	Account   TypePrefix = "t2_"
	Link      TypePrefix = "t3_"
	Message   TypePrefix = "t4_"
	Subreddit TypePrefix = "t5_"
	Award     TypePrefix = "t6_"
)

const logDivider = "--------------------------------------------------------------------"

type LimbRetrievalBot interface {
	BotLogger
	Clock
	//Api
}

type Clock interface {
	NowUTC() time.Time
}

type clock struct{}

func (c *clock) NowUTC() time.Time {
	return time.Now().UTC()
}

type Api interface {
	CheckParentId(parentId string)
	LogInteraction()
}

type api struct {
	//reddit.Bot
}

type limbRetrievalBot struct {
	BotLogger
	Clock
	//Api
}

func NewLimbRetrievalBot(logPath string) (LimbRetrievalBot, error) {
	logger := NewBotLogger(logPath)
	logPathErr := logger.initDir(logPath + eventDir)
	if logPathErr != nil {
		return nil, logPathErr
	}
	errPathErr := logger.initDir(logPath + errDir)
	if errPathErr != nil {
		return nil, errPathErr
	}

	return limbRetrievalBot{
		BotLogger: logger,
		Clock:     &clock{},
	}, nil
}

func RetrieveLimbs(comment *reddit.Comment, bot LimbRetrievalBot) string {
	var shrug = CheckContainsShrug(comment.Permalink, comment.Body, comment.BodyHTML, bot)

	if shrug != NoShrug {
		event := BotEvent{
			DateTime:     bot.NowUTC(),
			RedditorName: comment.Author,
			Subreddit:    comment.Subreddit,
			Event:        shrug.matchType(),
			Permalink:    comment.Permalink,
		}

		bot.LogEvent(&event)
		return shrug.commentResponse()
	}

	return ""
}

func CheckContainsShrug(permalink, body, bodyHtml string, b LimbRetrievalBot) Shrug {
	b.Debug(logDivider)
	b.Debugf("checking for shrug in comment at: %s", permalink)
	for _, s := range invalidShrugBodies {
		match, _ := regexp.Match(string(s), []byte(body))
		if match {
			b.Debugf("found a preliminary match against %s", s)

			// verify not inside code block
			isInsideCodeBlock, _ := regexp.Match(CodeBlockShrugPattern, []byte(bodyHtml))
			if !isInsideCodeBlock {
				b.Debug("match not inside code block, returning confirmed match")
				return s
			}
			b.Debug("preliminary match found inside code block, ignoring...")
		}
	}

	return NoShrug
}
