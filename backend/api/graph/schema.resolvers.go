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

func (r *mutationResolver) AddBannedSubreddit(ctx context.Context, input model.BannedSubredditInput) (*model.BannedSubreddit, error) {
	res, err := r.BannedSubredditHandler.Insert(input.Subreddit)
	if err != nil {
		return nil, err
	}

	return &model.BannedSubreddit{
		Subreddit:  res.Subreddit,
		InsertedAt: utils.FormatUnixToUTCString(res.InsertedAt),
	}, nil
}

func (r *mutationResolver) AddIgnoredUser(ctx context.Context, username string) (*model.IgnoredUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Thread(ctx context.Context, threadID string) (*model.Thread, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) BannedSubreddit(ctx context.Context, subreddit string) (*model.BannedSubreddit, error) {
	res, err := r.BannedSubredditHandler.Get(subreddit)

	// no result, not saved as banned subreddit
	if res.InsertedAt == 0 && err != nil {
		return &model.BannedSubreddit{
			Subreddit: subreddit,
		}, nil
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

func (r *queryResolver) BannedSubreddits(ctx context.Context) ([]*model.BannedSubreddit, error) {
	res, err := r.BannedSubredditHandler.GetAll()

	if err != nil {
		return []*model.BannedSubreddit{}, err
	}

	results := make([]*model.BannedSubreddit, len(*res))
	for i, subreddit := range *res {
		results[i] = &model.BannedSubreddit{
			Subreddit:  subreddit.Subreddit,
			InsertedAt: utils.FormatUnixToUTCString(subreddit.InsertedAt),
		}
	}

	return results, nil
}

func (r *queryResolver) IgnoredUser(ctx context.Context, username string) (*model.IgnoredUser, error) {
	return &model.IgnoredUser{
		Username:  "SomeRandomRedditor",
		IgnoredAt: utils.FormatUnixToUTCString(1653230094),
	}, nil
}

func (r *queryResolver) IgnoredUsers(ctx context.Context) ([]*model.IgnoredUser, error) {
	return []*model.IgnoredUser{
		{
			Username:  "SomeRandomRedditor",
			IgnoredAt: utils.FormatUnixToUTCString(1653230094),
		},
		{
			Username:  "SomeRandomRedditor2",
			IgnoredAt: utils.FormatUnixToUTCString(1653230094),
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
