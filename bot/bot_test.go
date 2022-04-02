package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/turnage/graw/reddit"
	"testing"
	"time"
)

const permalink = "https://reddit.com"
const subreddit = "programming"

var mockDate = time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)

type mockLimbBot struct {
	mock.Mock
}

func (m *mockLimbBot) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (m *mockLimbBot) Infof(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (m *mockLimbBot) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (m *mockLimbBot) Debugf(template string, args ...interface{}) {
	fmt.Printf(template+"\n", args...)
}

func (m *mockLimbBot) LogError(event *BotEvent) {
	m.Called(event)
}

func (m *mockLimbBot) LogEvent(event *BotEvent) {
	m.Called(event)
}

func (m *mockLimbBot) NowUTC() time.Time {
	return mockDate
}

func TestRetrieveLimbs(t *testing.T) {
	bot := new(mockLimbBot)
	bot.On("LogEvent", mock.Anything)
	bot.On("Info", mock.Anything)
	bot.On("Infof", mock.Anything)
	bot.On("Debug", mock.Anything)
	bot.On("Debugf", mock.Anything)

	testCases := []struct {
		comment        *reddit.Comment
		expectedString string
	}{
		{
			comment: &reddit.Comment{
				Body: "Some test string with no shrug",
			},
			expectedString: "",
		},
		{
			comment: &reddit.Comment{
				Body:      `¯\_(ツ)_/¯`,
				Author:    "BobbyTables",
				Permalink: permalink,
				Subreddit: subreddit,
			},
			expectedString: MissingLeftArmPattern.commentResponse(),
		},
		{
			comment: &reddit.Comment{
				Body:      `asdf¯\_(ツ)_/¯asdf`,
				Author:    "BobbyTables",
				Permalink: permalink,
				Subreddit: subreddit,
			},
			expectedString: MissingLeftArmPattern.commentResponse(),
		},
		{
			comment: &reddit.Comment{
				Body:      `¯\\_(ツ)_/¯`,
				Author:    "BobbyTables",
				Permalink: permalink,
				Subreddit: subreddit,
			},
			expectedString: MissingShouldersPattern.commentResponse(),
		},
		{
			comment: &reddit.Comment{
				Body:      `asdf¯\\_(ツ)_/¯asdf`,
				Author:    "BobbyTables",
				Permalink: permalink,
				Subreddit: subreddit,
			},
			expectedString: MissingShouldersPattern.commentResponse(),
		},
		// Verify HTML
		{
			comment: &reddit.Comment{
				Body:      "`¯\\_(ツ)_/¯`",
				BodyHTML:  `<div class="md"><p><code>¯\_(ツ)_/¯</code></p></div>`,
				Author:    "BobbyTables",
				Permalink: permalink,
				Subreddit: subreddit,
			},
			expectedString: NoShrug.commentResponse(),
		},
		{
			comment: &reddit.Comment{
				Body: "`¯\\_(ツ)_/¯`",
				BodyHTML: `<code>¯\_(ツ)_/¯
</code>`,
				Author:    "BobbyTables",
				Permalink: permalink,
				Subreddit: subreddit,
			},
			expectedString: NoShrug.commentResponse(),
		},
		{
			comment: &reddit.Comment{
				Body:      "`¯\\_(ツ)_/¯`",
				BodyHTML:  `<code>¯\_(ツ)_/¯\n</code>`,
				Author:    "BobbyTables",
				Permalink: permalink,
				Subreddit: subreddit,
			},
			expectedString: NoShrug.commentResponse(),
		},
	}

	for _, test := range testCases {
		result := RetrieveLimbs(test.comment, bot)
		assert.Equal(t, test.expectedString, result)
	}

}
