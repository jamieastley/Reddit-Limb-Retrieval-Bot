package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"api/graph/generated"
	"api/graph/model"
	"context"
	"fmt"
	"utils"
)

func (r *mutationResolver) AddBannedSubreddit(_ context.Context, input model.BannedSubredditInput) (*model.BannedSubreddit, error) {
	res, err := r.BannedSubredditHandler.Insert(input.Subreddit)
	if err != nil {
		return nil, err
	}

	return &model.BannedSubreddit{
		Subreddit:  res.Subreddit,
		InsertedAt: utils.FormatUnixToUTCString(res.InsertedAt),
	}, nil
}

func (r *mutationResolver) AddIgnoredUser(_ context.Context, username string) (*model.IgnoredUser, error) {
	res, err := r.IgnoredUserHandler.Insert(username)
	if err != nil {
		return nil, err
	}

	return &model.IgnoredUser{
		Username:  res.Username,
		IgnoredAt: utils.FormatUnixToUTCString(res.IgnoredAt),
	}, err
}

func (r *mutationResolver) RemoveIgnoredUser(_ context.Context, username string) (*model.DeleteMutationResponse, error) {
	res, err := r.IgnoredUserHandler.Remove(username)

	return &model.DeleteMutationResponse{
		AffectedRows: int(res),
	}, err
}

func (r *queryResolver) Thread(_ context.Context, threadID string) (*model.Thread, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BannedSubreddit(_ context.Context, subreddit string) (*model.BannedSubreddit, error) {
	res, err := r.BannedSubredditHandler.Get(subreddit)

	// no result, not saved as banned subreddit
	if res.InsertedAt == 0 && err != nil {
		return nil, nil
	}

	// some other error occurred, surface it
	if err != nil {
		return nil, err
	}

	return &model.BannedSubreddit{
		Subreddit:  res.Subreddit,
		InsertedAt: utils.FormatUnixToUTCString(res.InsertedAt),
	}, nil
}

func (r *queryResolver) BannedSubreddits(_ context.Context) ([]*model.BannedSubreddit, error) {
	res, err := r.BannedSubredditHandler.GetAll()

	if err != nil {
		return []*model.BannedSubreddit{}, err
	}

	results := make([]*model.BannedSubreddit, len(res))
	for i, subreddit := range res {
		results[i] = &model.BannedSubreddit{
			Subreddit:  subreddit.Subreddit,
			InsertedAt: utils.FormatUnixToUTCString(subreddit.InsertedAt),
		}
	}

	return results, nil
}

func (r *queryResolver) IgnoredUser(_ context.Context, username string) (*model.IgnoredUser, error) {
	user, err := r.IgnoredUserHandler.Get(username)

	if user.IgnoredAt == 0 && err == nil {
		return nil, nil
	}

	return &model.IgnoredUser{
		Username:  user.Username,
		IgnoredAt: utils.FormatUnixToUTCString(user.IgnoredAt),
	}, err
}

func (r *queryResolver) IgnoredUsers(_ context.Context) ([]*model.IgnoredUser, error) {
	results, err := r.IgnoredUserHandler.GetAll()

	if err != nil {
		return []*model.IgnoredUser{}, err
	}

	users := make([]*model.IgnoredUser, len(results))
	for i, user := range results {
		users[i] = &model.IgnoredUser{
			Username:  user.Username,
			IgnoredAt: utils.FormatUnixToUTCString(user.IgnoredAt),
		}
	}

	return users, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
