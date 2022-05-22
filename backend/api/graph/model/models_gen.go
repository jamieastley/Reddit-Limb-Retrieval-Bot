// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// Defines a banned subreddit.
//
// Subreddits are typically saved as banned when the bot fails to respond to comments within a given subreddit.
type BannedSubreddit struct {
	// The name of the subreddit.
	Subreddit string `json:"subreddit"`
	// The UTC ISO-8601 String of when the subreddit was saved as banned.
	InsertedAt string `json:"insertedAt"`
}

type BannedSubredditInput struct {
	// The name of the subreddit.
	//
	// Should exclude the leading 'r/', as this should be up to the caller/client to append if they so wish.
	Subreddit string `json:"subreddit"`
}

type Thread struct {
	ID         string `json:"id"`
	ThreadID   string `json:"threadId"`
	InsertedAt string `json:"insertedAt"`
}
