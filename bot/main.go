package main

import (
	"fmt"
	"github.com/turnage/graw"
	"github.com/turnage/graw/reddit"
	"os"
)

const botName = "LimbRetrieval-Bot"

var agentPath = os.Getenv("AGENT_PATH")
var loggingDirPath = os.Getenv("LOGGING_PATH")

type grawBot struct {
	reddit.Bot
	LimbRetrievalBot
}

func main() {

	if loggingDirPath == "" {
		loggingDirPath = "bot_logs/"
	}

	bot, err := reddit.NewBotFromAgentFile(agentPath, 0)
	if err != nil {
		fmt.Printf("Failed to init agent: %s", err)
	}
	cfg := graw.Config{
		SubredditComments: []string{
			"LimbRetrievalBotTest",
			"LockedLRBTest",
		},
	}

	lrb, err := NewLimbRetrievalBot(loggingDirPath)
	if err != nil {
		fmt.Printf("failed to initialise %s: %s", botName, err)
		os.Exit(1)
	}
	handler := &grawBot{bot, lrb}

	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		fmt.Printf("Failed to start graw run: %s", err)
	} else {
		// wait() returns an error message if the bot fails
		// TODO: handle reddit server issues
		fmt.Println("graw run failed: ", wait())
	}
}

func (b *grawBot) Comment(comment *reddit.Comment) error {

	response := RetrieveLimbs(comment, b.LimbRetrievalBot)

	if response != "" {
		if err := b.Reply(comment.Name, FormatLimbResponse(response)); err != nil {
			b.LogError(
				&BotEvent{
					DateTime:     b.NowUTC(),
					RedditorName: comment.Author,
					Subreddit:    comment.Subreddit,
					Event:        err.Error(),
					Permalink:    comment.Permalink,
				},
			)
		}
	}

	return nil
}
