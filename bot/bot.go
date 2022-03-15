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

func NewLimbRetrievalBot(filePath string) LimbRetrievalBot {
	return limbRetrievalBot{
		Logger: NewBotLogger(filePath),
		Clock:  &clock{},
	}
}

func RetrieveLimbs(comment *reddit.Comment, limbRetrievalBot LimbRetrievalBot) string {
	var shrug = checkContainsShrug(comment.Body, comment.BodyHTML)
	event := BotEvent{
		DateTime:     limbRetrievalBot.NowUTC(),
		RedditorName: comment.Author,
		Subreddit:    comment.Subreddit,
		Event:        shrug.matchType(),
		Permalink:    comment.Permalink,
	}

	if shrug != NoShrug {
		limbRetrievalBot.Log(&event)
		return shrug.commentResponse()
	}

	return ""
}

func checkContainsShrug(body string, bodyHtml string) Shrug {
	for _, s := range invalidShrugBodies {
		match, _ := regexp.Match(string(s), []byte(body))
		if match {
			fmt.Sprintf("found a match against %s", s)
		}

		if match {
			// TODO: first match passed, now check if within code block
			return s
		}
	}

	return NoShrug
}

func HandleError(err error) {

}
