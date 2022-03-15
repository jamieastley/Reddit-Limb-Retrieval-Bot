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
		SubredditComments: []string{"LimbRetrievalBotTest", "LockedLRBTest"},
	}
	handler := &grawBot{bot, NewLimbRetrievalBot(loggingDirPath)}
	if _, wait, err := graw.Run(handler, bot, cfg); err != nil {
		_ = fmt.Sprintf("Failed to start graw run: %s", err)
	} else {
		fmt.Println("graw run failed: ", wait())
	}
}

func (b *grawBot) Comment(comment *reddit.Comment) error {

	msg := RetrieveLimbs(comment, b.LimbRetrievalBot)

	if msg != "" {
		if err := b.Reply(comment.Name, msg); err != nil {
			HandleError(err)
		}
	}

	return nil
}
