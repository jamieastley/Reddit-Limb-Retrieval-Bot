package main

import (
	"fmt"
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

type LimbRetrievalBot interface {
	Logger
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
	Logger
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
		Logger: logger,
		Clock:  &clock{},
	}, nil
}

func RetrieveLimbs(comment *reddit.Comment, bot LimbRetrievalBot) string {
	var shrug = checkContainsShrug(comment.Body, comment.BodyHTML)

	if shrug != NoShrug {
		event := BotEvent{
			DateTime:     bot.NowUTC(),
			RedditorName: comment.Author,
			Subreddit:    comment.Subreddit,
			Event:        shrug.matchType(),
			Permalink:    comment.Permalink,
		}

		bot.Log(&event)
		return shrug.commentResponse()
	}

	return ""
}

func checkContainsShrug(body string, bodyHtml string) Shrug {
	for _, s := range invalidShrugBodies {
		match, _ := regexp.Match(string(s), []byte(body))
		if match {
			fmt.Printf("found a preliminary match against %s\n", s)

			// verify not inside code block, eg:
			// <div class="md"><p><code>¯\_(ツ)_/¯</code></p>
			isInsideCodeBlock, _ := regexp.Match(LiteralCodeShrugPattern, []byte(bodyHtml))
			if !isInsideCodeBlock {
				fmt.Println("match not inside code block, returning match")
				return s
			}
			fmt.Println("match not within code block, returning match")
		}
	}

	return NoShrug
}
