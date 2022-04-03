package main

import (
	"fmt"
	"github.com/turnage/graw/reddit"
	"regexp"
	"time"
)

const (
	logDivider     = "--------------------------------------------------------------------"
	commentDivider = "\n***"
	howToFooter    = "\n^^To&#32;prevent&#32;any&#32;more&#32;lost&#32;limbs&#32;throughout&#32;Reddit,&#32;correctly&#32;escape&#32;the&#32;arms&#32;and&#32;shoulders&#32;by&#32;typing&#32;the&#32;shrug&#32;as&#32;`¯\\\\\\_(ツ)_/¯`&#32;or&#32;`¯\\\\\\_(ツ)\\_/¯`"
	linkFooter     = "\n\n[^^Click&#32;here&#32;to&#32;see&#32;why&#32;this&#32;is&#32;necessary](https://np.reddit.com/r/OutOfTheLoop/comments/3fbrg3/is_there_a_reason_why_the_arm_is_always_missing/ctn5gbf/)"
)

type LimbRetrievalBot interface {
	BotLogger
	Clock
}

type Clock interface {
	NowUTC() time.Time
}

type clock struct{}

func (c *clock) NowUTC() time.Time {
	return time.Now().UTC()
}

type limbRetrievalBot struct {
	BotLogger
	Clock
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

	b.Debug("no matches found, ignoring comment...")
	return NoShrug
}

func FormatLimbResponse(msg string) string {
	return fmt.Sprintf("%s%s%s%s", msg, commentDivider, howToFooter, linkFooter)
}
