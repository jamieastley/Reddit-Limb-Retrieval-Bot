package main

import (
	"api/graph"
	"api/graph/generated"
	"fmt"
	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/jamieastley/limbretrievalbot/repository"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
	"time"
)

var mockDate = time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC)

type MockBannedSubreddits struct {
	sub *repository.BannedSubreddit
	err error
}

func (m *MockBannedSubreddits) Insert(subreddit string) (*repository.BannedSubreddit, error) {
	fmt.Printf("Inserting subreddit: %s", subreddit)
	return m.sub, m.err
}

func (m *MockBannedSubreddits) Get(subreddit string) (*repository.BannedSubreddit, error) {
	fmt.Printf("Retrieving subreddit: %s", subreddit)
	return m.sub, m.err
}

func TestBannedSubreddit(t *testing.T) {

	sub := &repository.BannedSubreddit{
		ID:         0,
		Subreddit:  "golang",
		InsertedAt: mockDate,
	}

	c := client.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			BannedSubreddit: &MockBannedSubreddits{
				sub: sub,
				err: nil,
			},
		},
	})))
	var resp struct {
		IsBannedSubreddit struct {
			ID         string
			Subreddit  string
			InsertedAt string
		}
	}
	c.MustPost(`{ 
		IsBannedSubreddit(subreddit:"golang") { 
			id 
			subreddit 
			insertedAt 
		} 
	}`, &resp)
	assert.Equal(t, strconv.Itoa(int(sub.ID)), resp.IsBannedSubreddit.ID)
	assert.Equal(t, sub.Subreddit, resp.IsBannedSubreddit.Subreddit)
	assert.Equal(t, sub.InsertedAt.UTC().Format(time.RFC3339), resp.IsBannedSubreddit.InsertedAt)
}
