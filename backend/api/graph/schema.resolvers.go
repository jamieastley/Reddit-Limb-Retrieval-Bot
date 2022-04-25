package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"context"
	"fmt"
	"strconv"
	"time"
)

func (r *mutationResolver) AddBannedSubreddit(ctx context.Context, input model.BannedSubredditInput) (*model.BannedSubreddit, error) {
	res, err := r.BannedSubreddit.Insert(input.Subreddit)
	if err != nil {
		return nil, err
	}

	return &model.BannedSubreddit{
		ID:         strconv.Itoa(int(res.ID)),
		Subreddit:  res.Subreddit,
		InsertedAt: res.InsertedAt.Format(time.RFC3339),
	}, nil
}

func (r *queryResolver) Thread(ctx context.Context, threadID string) (*model.Thread, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) IsBannedSubreddit(ctx context.Context, subreddit string) (*model.BannedSubreddit, error) {
	res, err := r.BannedSubreddit.Get(subreddit)
	if err != nil {
		return nil, err
	}

	return &model.BannedSubreddit{
		ID:         strconv.Itoa(int(res.ID)),
		Subreddit:  res.Subreddit,
		InsertedAt: res.InsertedAt.Format(time.RFC3339),
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
